package fileapi

import (
	"net/http"

	"github.com/go-chi/chi"
)

type FileAPI struct {
	Repos repos
}

func New(r repos) *FileAPI {
	a := &FileAPI{
		Repos: r,
	}
	return a
}

func (a *FileAPI) GetImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := chi.URLParam(r, "slug")

	img, err := a.Repos.GetImageBySlug(ctx, slug)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	// TOOD: Check if it is photo.
	w.Header().Set("Content-Type", img.MIME)
	w.Write(img.Img)
}
