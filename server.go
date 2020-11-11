package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GlitchyGlitch/typinger/graphql"
	// "github.com/GlitchyGlitch/typinger/graphql/dataloaders"
	"github.com/GlitchyGlitch/typinger/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pg"
)

const defaultPort = "8080"

func main() {
	opt, err := pg.ParseURL(os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		panic(err)
	}

	DB := postgres.New(opt)
	defer DB.Close()
	DB.AddQueryHook(&postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UserRepo{DB: DB}
	articleRepo := postgres.ArticleRepo{DB: DB}
	settingRepo := postgres.SettingRepo{DB: DB}

	schema := graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{UserRepo: userRepo, ArticleRepo: articleRepo, SettingRepo: settingRepo}})

	srv := handler.NewDefaultServer(schema)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
