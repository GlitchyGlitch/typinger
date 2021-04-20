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
