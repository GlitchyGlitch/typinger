package main

import (
	"context"
	"fmt"
	"time"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/GlitchyGlitch/typinger/test"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

const tdPath = "postgres/fixtures"

type user struct {
	Typename graphql.String `graphql:"__typename"`
	ID       graphql.String
	Name     graphql.String
	Email    graphql.String
}

type article struct {
	Typename     graphql.String `graphql:"__typename"`
	ID           graphql.String
	Title        graphql.String
	Content      graphql.String
	ThumbnailURL graphql.String
}

type image struct {
	Typename graphql.String `graphql:"__typename"`
	ID       graphql.String
	Name     graphql.String
	Slug     graphql.String
	URL      graphql.String
	MIME     graphql.String
}

var articlesDataSet = []article{
	{
		Typename:     "Article",
		ID:           "82ba242e-e853-499f-8873-f271c53aca6a",
		Title:        "First article",
		Content:      "First content.",
		ThumbnailURL: "http://www.example.com/path/to/photo1",
	},
	{
		Typename:     "Article",
		ID:           "c3eec2ac-0fd5-41ce-829a-6f3dd74cd102",
		Title:        "Second article",
		Content:      "Second content.",
		ThumbnailURL: "http://www.example.com/path/to/photo2",
	},
	{
		Typename:     "Article",
		ID:           "d50f5d60-6f59-4605-96b8-a96b9e9b17ea",
		Title:        "Third article",
		Content:      "Third content.",
		ThumbnailURL: "http://www.example.com/path/to/photo3",
	},
}

var userDataSet = []user{
	{
		Typename: "User",
		ID:       "e5f1c9af-fa8a-4a58-9909-d887ddf7e961",
		Name:     "First User",
		Email:    "first@example.com",
	},
	{
		Typename: "User",
		ID:       "d1451907-e1ec-4291-ab14-a9a314b56b6a",
		Name:     "Second User",
		Email:    "second@example.com",
	},
	{
		Typename: "User",
		ID:       "0e38a4bd-87a0-447f-93fd-b904c9f7f303",
		Name:     "Third User",
		Email:    "third@example.com",
	},
}

var imageDataSet = []image{
	{
		Typename: "Image",
		ID:       "0cf191b8-b60c-4aec-b698-ca2b64a3d0f7",
		Name:     "First image",
		Slug:     "first-slug",
		URL:      "http://0.0.0.0/img/first-slug",
		MIME:     "image/svg+xml",
	},
	{
		Typename: "Image",
		ID:       "60b390de-cd44-4cc2-9fbb-fd9bcaa42819",
		Name:     "Second image",
		Slug:     "second-slug",
		URL:      "http://0.0.0.0/img/second-slug",
		MIME:     "image/webp",
	},
	{
		Typename: "Image",
		ID:       "72e641f1-09ac-4444-b980-98be375f8efd",
		Name:     "Third image",
		Slug:     "third-slug",
		URL:      "http://0.0.0.0/img/third-slug",
		MIME:     "image/webp",
	},
}

var forbiddenError = "Operation forbidden."

func setup(auth bool, conf *config.Config) *graphql.Client {
	test.RenewTestData(conf.DBURL, tdPath)

	go startServer(conf, 2)
	time.Sleep(100 * time.Millisecond) // Wait for server startup
	url := fmt.Sprintf("http://%s/graphql", conf.Addr())
	if !auth {
		return graphql.NewClient(url, nil)
	}

	loginClient := graphql.NewClient(url, nil)
	var loginMut struct {
		Login graphql.String `graphql:"login(input: {email: \"first@example.com\", password:\"first\"})"`
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
