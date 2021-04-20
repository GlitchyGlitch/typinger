package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GlitchyGlitch/typinger/auth"
	"github.com/GlitchyGlitch/typinger/config"
	"github.com/GlitchyGlitch/typinger/dataloaders"
	"github.com/GlitchyGlitch/typinger/fileapi"
	"github.com/GlitchyGlitch/typinger/graphql"
	"github.com/GlitchyGlitch/typinger/jwtcontroller"
	"github.com/GlitchyGlitch/typinger/postgres"
	"github.com/GlitchyGlitch/typinger/services"
	"github.com/GlitchyGlitch/typinger/validator"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-pg/pg"
)

func router(gqlHandler http.Handler, repos *services.Repos, fAPI *fileapi.FileAPI, config *config.Config, tc *jwtcontroller.JWTController) *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{config.Addr()},
		AllowCredentials: true,
	}))
	router.Use(auth.Middleware(tc, repos))
	router.Use(dataloaders.Middleware(repos))

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", gqlHandler)
	router.Get(fmt.Sprintf("/%s/{slug}", config.ImgDir), fAPI.GetImage)
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

func startServer(conf *config.Config, testTime int) {
	opt, err := pg.ParseURL(conf.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	db := postgres.New(opt)
	defer db.Close()
	// db.AddQueryHook(&postgres.dbLogger{})

	tc := jwtcontroller.New(conf)
	repos := services.NewRepos(db, tc)
	fAPI := fileapi.New(repos)
	errPresenter := graphql.ErrorPresenter() // TODO: Make it works
	v := validator.New()
	handler := graphql.Handler(repos, conf, v, errPresenter)
	router := router(handler, repos, fAPI, conf, tc)
	srv := httpServer(conf, router)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Printf("🚀 Server running on %s", conf.Addr())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	if testTime != 0 {
		time.Sleep(time.Duration(testTime) * time.Second)
		close(sigChan)
	}

	<-sigChan
	log.Println("Shutting down...")

	ctx, ctxCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer ctxCancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := config.New()
	startServer(c, 0)
}
