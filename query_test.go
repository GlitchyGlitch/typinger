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

	t.Run("Get all articles", func(t *testing.T) {
		var query struct {
			Articles []struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			}
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "Article", string(query.Articles[0].Typename))
		require.Equal(t, "82ba242e-e853-499f-8873-f271c53aca6a", string(query.Articles[0].ID))
		require.Equal(t, "Post about awsomeness of Go", string(query.Articles[0].Title))
		require.Equal(t, "Go is awsome.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo0.jpg", string(query.Articles[0].ThumbnailURL))

		require.Equal(t, "Article", string(query.Articles[1].Typename))
		require.Equal(t, "c3eec2ac-0fd5-41ce-829a-6f3dd74cd102", string(query.Articles[1].ID))
		require.Equal(t, "Very important article about programming", string(query.Articles[1].Title))
		require.Equal(t, "Very important contetnt of article.", string(query.Articles[1].Content))
		require.Equal(t, "http://www.example.com/path/to/photo1.jpg", string(query.Articles[1].ThumbnailURL))

		require.Equal(t, "Article", string(query.Articles[1].Typename))
		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[2].ID))
		require.Equal(t, "Lorem ipsum article", string(query.Articles[2].Title))
		require.Equal(t, "Lorem ipsum dolor sit amet.", string(query.Articles[2].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2.jpg", string(query.Articles[2].ThumbnailURL))
	})

	t.Run("Get articles with pagination", func(t *testing.T) {
		var query struct {
			Articles []struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"articles(first:1, offset:2)"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "Article", string(query.Articles[0].Typename))
		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[0].ID))
		require.Equal(t, "Lorem ipsum article", string(query.Articles[0].Title))
		require.Equal(t, "Lorem ipsum dolor sit amet.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2.jpg", string(query.Articles[0].ThumbnailURL))
	})

	t.Run("Get articles with filter", func(t *testing.T) {
		var query struct {
			Articles []struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"articles(filter: {title: \"lorem ipsum\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "Article", string(query.Articles[0].Typename))
		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[0].ID))
		require.Equal(t, "Lorem ipsum article", string(query.Articles[0].Title))
		require.Equal(t, "Lorem ipsum dolor sit amet.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2.jpg", string(query.Articles[0].ThumbnailURL))

	})

	t.Run("No matches to filter", func(t *testing.T) {
		var query struct {
			Articles []struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"articles(filter: {title: \"No matches here\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)
		require.Len(t, query.Articles, 0)
	})

	t.Run("Invalid attribute of filter", func(t *testing.T) {
		var query struct {
			Articles []struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"articles(filter: {title: \"PtYaKbEfujgg012vzU54xnIO8uNFyrWnzT1s8qp469ktBIuyGAMV5DkIteDogufej8rqyOIv73H9dakOfD7gPoMk850GngJA17MyolQ39VwOzV65XlktlOQPf17MJPW56tcSFYuwUmh4tgKZiwYQ8TwC5onHlKiZtCBMVEwdV9Seb6ZzOk0ccvXP0NVStib6eghVeGPOPPtpVplgXasr4cYAp12TWCHXoxjE4JxKZR0c7CqCqQpziJGehZfyEBf9l\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.EqualError(t, err, "Title field is invalid.")
	})
}

func TestQueryUsers(t *testing.T) {
	conf := config.New()
	conf.Port = "8003"
	c := setup(false, conf)
	defer teardown(conf)

	t.Run("Get users forbidden", func(t *testing.T) {
		var query struct {
			Users []struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"users"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Get users with filter forbidden", func(t *testing.T) {
		var query struct {
			Users []struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"users(filter: {name:\"Dennis\", email:\"ritchie\"})"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Get users with pagination forbidden", func(t *testing.T) {
		var query struct {
			Users []struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"users(first: 2, offset: 1)\"`
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

	t.Run("Get users", func(t *testing.T) {
		var query struct {
			Users []struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"users"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "User", string(query.Users[0].Typename))
		require.Equal(t, "0e38a4bd-87a0-447f-93fd-b904c9f7f303", string(query.Users[0].ID))
		require.Equal(t, "Brendan Eich", string(query.Users[0].Name))
		require.Equal(t, "eich@gmail.com", string(query.Users[0].Email))

		require.Equal(t, "User", string(query.Users[1].Typename))
		require.Equal(t, "d1451907-e1ec-4291-ab14-a9a314b56b6a", string(query.Users[1].ID))
		require.Equal(t, "Guido van Rossum", string(query.Users[1].Name))
		require.Equal(t, "rossum@gmail.com", string(query.Users[1].Email))

		require.Equal(t, "User", string(query.Users[2].Typename))
		require.Equal(t, "e5f1c9af-fa8a-4a58-9909-d887ddf7e961", string(query.Users[2].ID))
		require.Equal(t, "Dennis Ritchie", string(query.Users[2].Name))
		require.Equal(t, "ritchie@gmail.com", string(query.Users[2].Email))

	})

	t.Run("Get users with filter", func(t *testing.T) {
		var query struct {
			Users []struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"users(filter:{name:\"Dennis\", email:\"ritchie\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "User", string(query.Users[0].Typename))
		require.Equal(t, "e5f1c9af-fa8a-4a58-9909-d887ddf7e961", string(query.Users[0].ID))
		require.Equal(t, "Dennis Ritchie", string(query.Users[0].Name))
		require.Equal(t, "ritchie@gmail.com", string(query.Users[0].Email))
	})

	t.Run("Get users with pagination", func(t *testing.T) {
		var query struct {
			Users []struct {
				Typename graphql.String `graphql:"__typename"`
				ID       graphql.String
				Name     graphql.String
				Email    graphql.String
			} `graphql:"users(first:1, offset:1)"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "User", string(query.Users[0].Typename))
		require.Equal(t, "d1451907-e1ec-4291-ab14-a9a314b56b6a", string(query.Users[0].ID))
		require.Equal(t, "Guido van Rossum", string(query.Users[0].Name))
		require.Equal(t, "rossum@gmail.com", string(query.Users[0].Email))
	})
}
