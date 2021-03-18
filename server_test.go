package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/GlitchyGlitch/typinger/test"
	"github.com/shurcooL/graphql"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

const tdPath = "postgres/test_data"

func setup(auth bool, conf *config.Config) *graphql.Client {
	test.RenewTestData(conf.DBURL, tdPath)

	go startServer(conf)
	time.Sleep(100 * time.Millisecond) // Wait for server startup
	url := fmt.Sprintf("http://%s/graphql", conf.Addr())
	if !auth {
		return graphql.NewClient(url, nil)
	}

	loginClient := graphql.NewClient(url, nil)
	var loginMut struct {
		Login graphql.String `graphql:"login(input: {email: \"ritchie@gmail.com\", password:\"ritchie\"})"`
	}
	err := loginClient.Mutate(context.Background(), &loginMut, nil)
	if err != nil {
		panic(err)
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(loginMut.Login)},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	// Setup authenticated client
	return graphql.NewClient(url, httpClient)
}
func teardown(conf *config.Config) {
	test.MigrateTestData("down", conf.DBURL, tdPath)
}

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

	t.Run("Get specified page of articles", func(t *testing.T) {
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

	t.Run("Get articles specified by filter", func(t *testing.T) {
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

func TestMutationArticles(t *testing.T) {
	conf := config.New()
	conf.Port = "8001"
	c := setup(false, conf)
	defer teardown(conf)

	t.Run("Creating artcile forbidden", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Artcile created as test\", content:\"Nothing important here.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Updateing article forbidden", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Artcile created as test\", content:\"Nothing important here.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.EqualError(t, err, "Operation forbidden.")
	})

	t.Run("Deleting artcile forbidden", func(t *testing.T) {
		var mutation struct {
			Deleted bool `graphql:"deleteArticle(id:\"d50f5d60-6f59-4605-96b8-a96b9e9b17ea\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.EqualError(t, err, "Operation forbidden.")
		require.False(t, mutation.Deleted)
	})
}

func TestMutationArticlesAuthenticated(t *testing.T) {
	conf := config.New()
	conf.Port = "8002"
	c := setup(true, conf)
	defer teardown(conf)

	t.Run("Create article", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Artcile created as test\", content:\"Nothing important here.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.NoError(t, err)
		require.Equal(t, "Article", string(mutation.Article.Typename))
		require.True(t, test.IsValidUUID(string(mutation.Article.ID)))
		require.Equal(t, "Artcile created as test", string(mutation.Article.Title))
		require.Equal(t, "Nothing important here.", string(mutation.Article.Content))
		require.Equal(t, "http://www.example.com/path/to/photo3.jpg", string(mutation.Article.ThumbnailURL))
	})

	t.Run("Create two same articles", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Article same as the other\", content:\"Nothing important here.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)

		err = c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Resource already exists.")
	})

	t.Run("Create article with no title", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"\", content:\"Lorem ipsum dolor sit amet.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Title field is invalid.")
	})

	t.Run("Create article with no content", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Example article\", content:\"\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Content field is invalid.")
	})

	t.Run("Create article with invalid thumbnailUrl", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Example article\", content:\"Lorem ipsum dolor sit amet.\", thumbnailUrl:\"Just a text\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Thumbnail URL field is invalid.")
	})

	t.Run("Create article with no thumbnailUrl", func(t *testing.T) {
		var mutation struct {
			Article struct {
				Typename     graphql.String `graphql:"__typename"`
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Example article\", content:\"Lorem ipsum dolor sit amet.\", thumbnailUrl:\"Just a text\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "Thumbnail URL field is invalid.")
	})

	t.Run("Delete article", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteArticle(id:\"82ba242e-e853-499f-8873-f271c53aca6a\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)
		require.True(t, bool(mutation.Deleted))
	})

	t.Run("Delete article not found", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteArticle(id:\"e5f1c9af-fa8a-4a58-9909-d887ddf7e961\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "No data found.")
		require.False(t, bool(mutation.Deleted))
	})

	t.Run("Delete article invalid id", func(t *testing.T) {
		var mutation struct {
			Deleted graphql.Boolean `graphql:"deleteArticle(id:\"NotAUUID\")"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.EqualError(t, err, "ID field is invalid.")
		require.False(t, bool(mutation.Deleted))
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
			} `graphql:"users(filter: {name:\"Dennis\", email:\"ritchie\"})\"`
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
			} `graphql:"users(filter: {name:\"Dennis\", email:\"ritchie\"})\"`
		}
		err := c.Query(context.Background(), &query, nil)
		require.EqualError(t, err, "Operation forbidden.")
	})
}
