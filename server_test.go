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

func TestQueryArticles(t *testing.T) { //TODO: chceck typenames
	conf := config.New()
	tdPath := "postgres/test_data"

	test.RenewTestData(conf.DBURL, tdPath)
	defer test.MigrateTestData("down", conf.DBURL, tdPath)
	q := startServer(conf)
	defer close(q)
	time.Sleep(100 * time.Millisecond) // Wait for server startup

	c := graphql.NewClient(fmt.Sprintf("http://%s/graphql", conf.Addr()), nil)

	t.Run("Get all articles", func(t *testing.T) {
		var query struct {
			Articles []struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			}
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "82ba242e-e853-499f-8873-f271c53aca6a", string(query.Articles[0].ID))
		require.Equal(t, "Post about awsomeness of Go", string(query.Articles[0].Title))
		require.Equal(t, "Go is awsome.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo0.jpg", string(query.Articles[0].ThumbnailURL))

		require.Equal(t, "c3eec2ac-0fd5-41ce-829a-6f3dd74cd102", string(query.Articles[1].ID))
		require.Equal(t, "Very important article about programming", string(query.Articles[1].Title))
		require.Equal(t, "Very important contetnt of article.", string(query.Articles[1].Content))
		require.Equal(t, "http://www.example.com/path/to/photo1.jpg", string(query.Articles[1].ThumbnailURL))

		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[2].ID))
		require.Equal(t, "Lorem ipsum article", string(query.Articles[2].Title))
		require.Equal(t, "Lorem ipsum dolor sit amet.", string(query.Articles[2].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2.jpg", string(query.Articles[2].ThumbnailURL))
	})

	t.Run("Get specified page of articles", func(t *testing.T) {
		var query struct {
			Articles []struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"articles(first:1, offset:2)"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)

		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[0].ID))
		require.Equal(t, "Lorem ipsum article", string(query.Articles[0].Title))
		require.Equal(t, "Lorem ipsum dolor sit amet.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2.jpg", string(query.Articles[0].ThumbnailURL))
	})

	t.Run("Get articles specified by filter", func(t *testing.T) {
		var query struct {
			Articles []struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"articles(filter: {title: \"lorem ipsum\"})"`
		}
		err := c.Query(context.Background(), &query, nil)

		require.NoError(t, err)
		require.Equal(t, "d50f5d60-6f59-4605-96b8-a96b9e9b17ea", string(query.Articles[0].ID))
		require.Equal(t, "Lorem ipsum article", string(query.Articles[0].Title))
		require.Equal(t, "Lorem ipsum dolor sit amet.", string(query.Articles[0].Content))
		require.Equal(t, "http://www.example.com/path/to/photo2.jpg", string(query.Articles[0].ThumbnailURL))

	})

	t.Run("No matches to filter", func(t *testing.T) {
		var query struct {
			Articles []struct {
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

}

func TestMutationArticlesAuthenticated(t *testing.T) {
	conf := config.New()
	tdPath := "postgres/test_data"

	test.RenewTestData(conf.DBURL, tdPath)
	defer test.MigrateTestData("down", conf.DBURL, tdPath)

	q := startServer(conf)
	defer close(q)
	time.Sleep(100 * time.Millisecond) // Wait for server startup

	// Prepare http client for graphql client
	loginClient := graphql.NewClient(fmt.Sprintf("http://%s/graphql", conf.Addr()), nil)
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
	c := graphql.NewClient(fmt.Sprintf("http://%s/graphql", conf.Addr()), httpClient)

	t.Run("Create article", func(t *testing.T) {
		var mutation struct {
			Article struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Artcile created as test\", content:\"Nothing important here.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.NoError(t, err)
		require.True(t, test.IsValidUUID(string(mutation.Article.ID)))
		require.Equal(t, "Artcile created as test", string(mutation.Article.Title))
		require.Equal(t, "Nothing important here.", string(mutation.Article.Content))
		require.Equal(t, "http://www.example.com/path/to/photo3.jpg", string(mutation.Article.ThumbnailURL))
	})

	t.Run("Create two same articles", func(t *testing.T) {
		var mutation struct {
			Article struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Article same as the other\", content:\"Nothing important here.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.NoError(t, err)

		err = c.Mutate(context.Background(), &mutation, nil)
		require.Equal(t, "Resource already exists.", err.Error())
	})

	t.Run("Create article with no title", func(t *testing.T) {
		var mutation struct {
			Article struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"\", content:\"Lorem ipsum dolor sit amet.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.Equal(t, "Title field is invalid.", err.Error())
	})

	t.Run("Create article with no content", func(t *testing.T) {
		var mutation struct {
			Article struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Example article\", content:\"\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.Equal(t, "Content field is invalid.", err.Error())
	})

	t.Run("Create article with invalid thumbnailUrl", func(t *testing.T) {
		var mutation struct {
			Article struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Example article\", content:\"Lorem ipsum dolor sit amet.\", thumbnailUrl:\"Just a text\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.Equal(t, "Thumbnail URL field is invalid.", err.Error())
	})

	t.Run("Create article with no thumbnailUrl", func(t *testing.T) {
		var mutation struct {
			Article struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Example article\", content:\"Lorem ipsum dolor sit amet.\", thumbnailUrl:\"Just a text\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)
		require.Equal(t, "Thumbnail URL field is invalid.", err.Error())
	})
}

func TestMutationArticlesUnauthenticated(t *testing.T) {
	conf := config.New()

	q := startServer(conf)
	defer close(q)
	time.Sleep(100 * time.Millisecond) // Wait for server startup

	c := graphql.NewClient(fmt.Sprintf("http://%s/graphql", conf.Addr()), nil)

	t.Run("Forbid creating article", func(t *testing.T) {
		var mutation struct {
			Article struct {
				ID           graphql.String
				Title        graphql.String
				Content      graphql.String
				ThumbnailURL graphql.String
			} `graphql:"createArticle(input: {title:\"Artcile created as test\", content:\"Nothing important here.\", thumbnailUrl:\"http://www.example.com/path/to/photo3.jpg\"})"`
		}
		err := c.Mutate(context.Background(), &mutation, nil)

		require.Equal(t, "Operation forbidden.", err.Error())

	})
}
