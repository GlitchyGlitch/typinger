package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/dataloaders"
	"github.com/GlitchyGlitch/typinger/graphql"
	"github.com/GlitchyGlitch/typinger/postgres"
	"github.com/GlitchyGlitch/typinger/services"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-pg/pg"
)

const defaultPort = "8080"

func main() {
	opt, err := pg.ParseURL(os.Getenv("POSTGRES_URL")) //TODO: Move it to config struct
	if err != nil {
		panic(err)
	}

	DB := postgres.New(opt)
	defer DB.Close()
	// DB.AddQueryHook(&postgres.DBLogger{})

	port := os.Getenv("PORT") //TODO: Move it to config struct
	if port == "" {
		port = defaultPort
	}

	repos := services.NewRepos(DB)
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"localhost:8080"}, //TODO: move it to config struct (loaded by variable)
		AllowCredentials: true,
	}))
	router.Use(auth.Middleware(repos))
	router.Use(dataloaders.Middleware(repos))

	errPresenter := graphql.ErrorPresenter()
	s := graphql.Server(repos, errPresenter)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", s)

	log.Printf("🚀 Server running on http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
