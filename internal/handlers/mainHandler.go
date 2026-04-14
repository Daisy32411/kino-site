package handlers

import (
	"bytes"
	"database/sql"
	"html/template"
	"kino-site/internal/models"
	"kino-site/internal/storage/postgres"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	MovieStorage *postgres.MovieStorage
	UserStorage  *postgres.UserStorage
	Templates	 *template.Template
}

func (h *Handler) MainHandler(w http.ResponseWriter, r *http.Request) {
	var email string
	cookie, err := r.Cookie("user_email")
	if err == nil && cookie.Value != "" {
		email = cookie.Value
	}

	query := r.URL.Query().Get("q")

	var movies []models.Movie
	var errDB error
	if query != "" {
		movies, errDB = h.MovieStorage.Search(query)
	} else {
		movies, errDB = h.MovieStorage.GetAll()
	}
	if errDB != nil {
		http.Error(w, "DB error", 500)
		return
	}

	data := struct {
		Title  string
		Movies []models.Movie
		Email  string
	}{}

	if query != "" {
		data.Title = "Фильмы по поиску: " + query
	} else {
		data.Title = "Все фильмы"
	}
	data.Movies = movies
	data.Email = email

	err = h.Templates.ExecuteTemplate(w, "main", data)
	if err != nil {
		log.Println("Failed to execute template", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) MoviePage(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	movie, err := h.MovieStorage.GetBySlug(slug)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Movie not found", 404)
			return
		}
		log.Println("DB error:", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	var email string
	if cookie, err := r.Cookie("user_email"); err == nil {
		email = cookie.Value
	}

	data := struct {
		Movie *models.Movie
		Email string
	}{
		Movie: movie,
		Email: email,
	}

	buf := new(bytes.Buffer)
	if err := h.Templates.ExecuteTemplate(buf, "movie", data); err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
}

func (h *Handler) MovieByGenre(w http.ResponseWriter, r *http.Request) {
	genre := chi.URLParam(r, "genre")
	decoded, _ := url.QueryUnescape(genre)

	movies, err := h.MovieStorage.GetByGenre(decoded)
	if err != nil {
		http.Error(w, "Movies not found", 404)
		return
	}

	var email string
	if cookie, err := r.Cookie("user_email"); err == nil {
		email = cookie.Value
	}

	data := struct {
		Title  string
		Movies []models.Movie
		Email  string
	}{
		Title:  "Фильмы жанра: " + decoded,
		Movies: movies,
		Email:  email,
	}

	buf := new(bytes.Buffer)
	if err := h.Templates.ExecuteTemplate(buf, "main", data); err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
}

func (h *Handler) MoviesByDirector(w http.ResponseWriter, r *http.Request) {
	director := chi.URLParam(r, "director")
	movies, err := h.MovieStorage.GetByDirector(director)
	if err != nil {
		http.Error(w, "Movies not found", 404)
		return
	}

	var email string
	if cookie, err := r.Cookie("user_email"); err == nil {
		email = cookie.Value
	}

	data := struct {
		Title  string
		Movies []models.Movie
		Email  string
	}{
		Title:  "Фильмы режисёра: " + director,
		Movies: movies,
		Email:  email,
	}

	buf := new(bytes.Buffer)
	if err := h.Templates.ExecuteTemplate(buf, "main", data); err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
}
