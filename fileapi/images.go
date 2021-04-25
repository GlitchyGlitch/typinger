package fileapi

import (
	"net/http"

	"github.com/GlitchyGlitch/typinger/config"
	"github.com/go-chi/chi"
)

type DBFileAPI struct { //TODO: There probably is cleaner way to implement this struct
	Repos  repos
	Config *config.Config
}

func New(r repos, c *config.Config) *DBFileAPI {
	a := &DBFileAPI{
		Repos:  r,
		Config: c,
	}
	return a
}

func (a *DBFileAPI) ImageHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		slug := chi.URLParam(r, "slug")

		img, err := a.Repos.GetImageBySlug(ctx, slug)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		// TODO: Check if it is photo.
		w.Header().Set("Content-Type", img.MIME)
		w.Write(img.Img)
	}
}
