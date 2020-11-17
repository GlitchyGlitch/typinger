package main

import (
	"log"
	"net/http"

	"os"

	"github.com/GlitchyGlitch/typinger/dataloaders"
	"github.com/GlitchyGlitch/typinger/graphql"
	"github.com/GlitchyGlitch/typinger/postgres"

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

	repos := postgres.NewRepos(DB)
	dl := dataloaders.NewRetriever()
	dlMiddleware := dataloaders.Middleware(repos)
	handler := graphql.Handler(repos, dl)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dlMiddleware(handler))

	log.Printf("🚀 Server running on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
