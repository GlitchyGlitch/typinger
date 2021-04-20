package main

import (
	"context"
	"testing"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/shurcooL/graphql"
	"github.com/stretchr/testify/require"
)

func TestQueryArticles(t *testing.T) {
	conf := config.New()
	conf.Port = "8080"
	c := setup(false, conf)
	defer teardown(conf)

	type article struct {
		Typename     graphql.String `graphql:"__typename"`
		ID           graphql.String
		Title        graphql.String
		Content      graphql.String
		ThumbnailURL graphql.String
	}

	t.Run("Get all articles", func(t *testing.T) {
		var query struct {
			Articles []article
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "Article", string(query.Articles[0].Typename))
		require.Equal(t, "82ba242e-e853-499f-8873-f271c53aca6a", string(query.Articles[0].ID))
		require.Equal(t, "First article", string(query.Articles[0].Title))
		require.Equal(t, "First content.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo1", string(query.Articles[0].ThumbnailURL))

		require.Equal(t, "Article", string(query.Articles[1].Typename))
		require.Equal(t, "c3eec2ac-0fd5-41ce-829a-6f3dd74cd102", string(query.Articles[1].ID))
		require.Equal(t, "Second article", string(query.Articles[1].Title))
		require.Equal(t, "Second content.", string(query.Articles[1].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2", string(query.Articles[1].ThumbnailURL))

		require.Equal(t, "Article", string(query.Articles[2].Typename))
		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[2].ID))
		require.Equal(t, "Third article", string(query.Articles[2].Title))
		require.Equal(t, "Third content.", string(query.Articles[2].Content))
		require.Equal(t, "http://www.example.com/path/to/photo3", string(query.Articles[2].ThumbnailURL))
	})

	t.Run("Get articles with pagination", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(first:1, offset:2)"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "Article", string(query.Articles[0].Typename))
		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[0].ID))
		require.Equal(t, "Third article", string(query.Articles[0].Title))
		require.Equal(t, "Third content.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo3", string(query.Articles[0].ThumbnailURL))
	})

	t.Run("Get articles with filter", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(filter: {title: \"sec\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)
		require.Equal(t, "Article", string(query.Articles[0].Typename))
		require.Equal(t, "c3eec2ac-0fd5-41ce-829a-6f3dd74cd102", string(query.Articles[0].ID))
		require.Equal(t, "Second article", string(query.Articles[0].Title))
		require.Equal(t, "Second content.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2", string(query.Articles[0].ThumbnailURL))

	})

	t.Run("No matches to filter", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(filter: {title: \"No matches here\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)
		require.Len(t, query.Articles, 0)
	})

	t.Run("Invalid attribute of filter", func(t *testing.T) {
		var query struct {
			Articles []article `graphql:"articles(filter: {title: \"PtYaKbEfujgg012vzU54xnIO8uNFyrWnzT1s8qp469ktBIuyGAMV5DkIteDogufej8rqyOIv73H9dakOfD7gPoMk850GngJA17MyolQ39VwOzV65XlktlOQPf17MJPW56tcSFYuwUmh4tgKZiwYQ8TwC5onHlKiZtCBMVEwdV9Seb6ZzOk0ccvXP0NVStib6eghVeGPOPPtpVplgXasr4cYAp12TWCHXoxjE4JxKZR0c7CqCqQpziJGehZfyEBf9l\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.EqualError(t, err, "Field \"title\" is invalid.")
	})
}

func TestQueryUsers(t *testing.T) {
	conf := config.New()
	conf.Port = "8003"
	c := setup(false, conf)
	defer teardown(conf)

	type user struct {
		Typename graphql.String `graphql:"__typename"`
		ID       graphql.String
		Name     graphql.String
		Email    graphql.String
	}

	t.Run("Get users forbidden", func(t *testing.T) {
		var query struct {
			Users []user
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Get users with filter forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter: {name:\"Dennis\", email:\"ritchie\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Get users with pagination forbidden", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first: 2, offset: 1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})
}

func TestQueryUsersAuthenticated(t *testing.T) {
	conf := config.New()
	conf.Port = "8004"
	c := setup(true, conf)
	defer teardown(conf)

	type user struct {
		Typename graphql.String `graphql:"__typename"`
		ID       graphql.String
		Name     graphql.String
		Email    graphql.String
	}

	t.Run("Get users", func(t *testing.T) {
		var query struct {
			Users []user
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)
		require.Equal(t, "User", string(query.Users[0].Typename))
		require.Equal(t, "e5f1c9af-fa8a-4a58-9909-d887ddf7e961", string(query.Users[0].ID))
		require.Equal(t, "First User", string(query.Users[0].Name))
		require.Equal(t, "first@example.com", string(query.Users[0].Email))

		require.Equal(t, "User", string(query.Users[1].Typename))
		require.Equal(t, "d1451907-e1ec-4291-ab14-a9a314b56b6a", string(query.Users[1].ID))
		require.Equal(t, "Second User", string(query.Users[1].Name))
		require.Equal(t, "second@example.com", string(query.Users[1].Email))

		require.Equal(t, "User", string(query.Users[2].Typename))
		require.Equal(t, "0e38a4bd-87a0-447f-93fd-b904c9f7f303", string(query.Users[2].ID))
		require.Equal(t, "Third User", string(query.Users[2].Name))
		require.Equal(t, "third@example.com", string(query.Users[2].Email))

	})

	t.Run("Get users with filter", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(filter:{name:\"second\", email:\"ex\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "User", string(query.Users[0].Typename))
		require.Equal(t, "d1451907-e1ec-4291-ab14-a9a314b56b6a", string(query.Users[0].ID))
		require.Equal(t, "Second User", string(query.Users[0].Name))
		require.Equal(t, "second@example.com", string(query.Users[0].Email))
	})

	t.Run("Get users with pagination", func(t *testing.T) {
		var query struct {
			Users []user `graphql:"users(first:1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "User", string(query.Users[0].Typename))
		require.Equal(t, "d1451907-e1ec-4291-ab14-a9a314b56b6a", string(query.Users[0].ID))
		require.Equal(t, "Second User", string(query.Users[0].Name))
		require.Equal(t, "second@example.com", string(query.Users[0].Email))
	})
}

func TestQueryImage(t *testing.T) {
	conf := config.New()
	conf.Port = "8007"
	c := setup(false, conf)
	defer teardown(conf)

	type image struct {
		Typename graphql.String `graphql:"__typename"`
		ID       graphql.String
		Name     graphql.String
		URL      graphql.String
		MIME     graphql.String
	}

	t.Run("Get images forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Get images with pagination forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(first:1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Get images with filter forbidden", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images(filter: {name:\"sec\", slug:\"slu\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})
}
func TestQueryImageAuthentication(t *testing.T) {
	conf := config.New()
	conf.Port = "8008"
	c := setup(true, conf)
	defer teardown(conf)

	type image struct {
		Typename graphql.String `graphql:"__typename"`
		ID       graphql.String
		Name     graphql.String
		URL      graphql.String
		MIME     graphql.String
	}

	t.Run("Get images", func(t *testing.T) {
		var query struct {
			Images []image `graphql:"images"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.NoError(t, err)

		expected := []image{
			{
				Typename: "Image",
				ID:       "0cf191b8-b60c-4aec-b698-ca2b64a3d0f7",
				Name:     "First image",
				URL:      "http://0.0.0.0/img/first-slug",
				MIME:     "image/svg+xml",
			},
			{
				Typename: "Image",
				ID:       "60b390de-cd44-4cc2-9fbb-fd9bcaa42819",
				Name:     "Second image",
				URL:      "http://0.0.0.0/img/second-slug",
				MIME:     "image/webp",
			},
			{
				Typename: "Image",
				ID:       "72e641f1-09ac-4444-b980-98be375f8efd",
				Name:     "Third image",
				URL:      "http://0.0.0.0/img/third-slug",
				MIME:     "image/webp",
			},
		}

		require.Equal(t, query.Images, expected)
	})
}
