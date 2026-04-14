package main

import (
	"html/template"
	"kino-site/internal/handlers"
	"kino-site/internal/storage/postgres"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := postgres.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := postgres.InitDB(db); err != nil {
		log.Fatal(err)
	}

	templates := template.Must(template.ParseFiles(
		"web/templates/main.html",
		"web/templates/movie.html",
		"web/templates/register.html",
		"web/templates/login.html",
		"web/templates/partials/header.html",
		"web/templates/dashboard.html",
		"web/templates/header_logged_in.html",
		"web/templates/profile.html",
	))

	movieStorage := &postgres.MovieStorage{DB: db}
	userStorage := &postgres.UserStorage{DB: db}
	h := &handlers.Handler{
		MovieStorage: movieStorage,
		UserStorage:  userStorage,
		Templates:    templates,
	}

	router := chi.NewRouter()
	
	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Главная страница
	router.Get("/", h.MainHandler)

	// Фильмы
	router.Get("/movie/{slug}", h.MoviePage)

	// Фильмы по жанру
	router.Get("/genre/{genre}", h.MovieByGenre)

	// Фильмы по директору
	router.Get("/director/{director}", h.MoviesByDirector)

	// Регистрация
	router.Get("/register", h.RegisterPage)
	router.Post("/register", h.RegisterPost)

	// Вход
	router.Get("/login", h.LoginPage)
	router.Post("/login", h.LoginPost)

	// Дашборд для авторизованных
	router.With(h.AuthMiddleware).Get("/dashboard", h.DashboardHandler)

	// Выход
	router.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "user_email",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Профиль
	router.Get("/profile", h.ProfilePage)
	router.Post("/logout", h.Logout)

	log.Println("Server started at :1111")
	http.ListenAndServe(":1111", router)
}