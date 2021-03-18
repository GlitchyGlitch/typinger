package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/config"
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

func router(handler http.Handler, repos *services.Repos, config *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.Addr()},
		AllowCredentials: true,
	}))
	router.Use(auth.Middleware(repos))
	router.Use(dataloaders.Middleware(repos))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", handler)
	return router
}

func httpServer(conf *config.Config, router *chi.Mux) *http.Server {
	srv := &http.Server{
		Addr:         conf.Addr(),
		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		IdleTimeout:  conf.IdleTimeout,
		Handler:      router,
	}
	return srv
}

func startServer(conf *config.Config) {
	opt, err := pg.ParseURL(conf.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	DB := postgres.New(opt)
	defer DB.Close()
	// DB.AddQueryHook(&postgres.DBLogger{})

	repos := services.NewRepos(DB)
	errPresenter := graphql.ErrorPresenter() // TODO: Make it works
	v := validator.New()
	handler := graphql.Handler(repos, v, errPresenter)
	router := router(handler, repos, conf)
	srv := httpServer(conf, router)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Printf("🚀 Server running on %s", conf.Addr())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	<-sigChan
	log.Println("Shutting down...")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxCancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	c := config.New()
	startServer(c)
}
