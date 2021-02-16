package main

import (
	"net/http"
	"os"

	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/dataloaders"
	"github.com/GlitchyGlitch/typinger/graphql"
	"github.com/GlitchyGlitch/typinger/postgres"
	"github.com/GlitchyGlitch/typinger/services"
	"github.com/GlitchyGlitch/typinger/validator"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-pg/pg"
)

const defaultPort = "8080"

func startServer(ip, port string) {
	opt, err := pg.ParseURL(os.Getenv("POSTGRES_URL")) //TODO: Move it to config struct
	if err != nil {
		panic(err)
	}

	DB := postgres.New(opt)
	defer DB.Close()
	// DB.AddQueryHook(&postgres.DBLogger{})

	repos := services.NewRepos(DB)
	errPresenter := graphql.ErrorPresenter()
	router := chi.NewRouter()
	valid := validator.New()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"localhost:8080"}, //TODO: move it to config struct (loaded by variable)
		AllowCredentials: true,
	}))
	router.Use(auth.Middleware(repos))
	router.Use(dataloaders.Middleware(repos))

	h := graphql.Handler(repos, valid, errPresenter)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/graphql", h)

	http.ListenAndServe(ip+":"+port, router)
}

func main() {
	startServer("localhost", "8080")
}
