package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RegisterPage(w http.ResponseWriter, r *http.Request) {
    err := h.Templates.ExecuteTemplate(w, "register", nil)
    if err != nil {
        http.Error(w, "Template execution error: " + err.Error(), 500)
        return
    }
}

func (h *Handler) RegisterPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error", 500)
		return
	}

	err = h.UserStorage.CreateUser(email, string(hash))
	if err != nil {
		http.Error(w, "User exists or DB error", 400)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (h *Handler) LoginPage(w http.ResponseWriter, r *http.Request) {
    err := h.Templates.ExecuteTemplate(w, "login", nil)
    if err != nil {
        http.Error(w, "Failed to render login page", 500)
    }
}

func (h *Handler) LoginPost(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.UserStorage.GetByEmail(email)
	if err != nil {
		http.Error(w, "User not found", 404)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid password", 400)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "user_email",
		Value: email,
		Path: "/",
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_email")
		if err != nil || cookie.Value == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return 
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user_email")
	email := cookie.Value

	data := struct {
		Title string
		Email string
	}{
		Title: "Добро пожаловать",
		Email: email,
	}

	err := h.Templates.ExecuteTemplate(w, "dashboard", data)
	if err != nil {
		http.Error(w, "Template execution error: " + err.Error(), 500)
		return
	}
}

func (h *Handler) ProfilePage(w http.ResponseWriter, r *http.Request) {
    email, _ := r.Cookie("user_email")
    if email == nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    data := struct{ Email string }{Email: email.Value}
    if err := h.Templates.ExecuteTemplate(w, "profile", data); err != nil {
        http.Error(w, "Internal server error", 500)
    }
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:   "user_email",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    })
    http.Redirect(w, r, "/", http.StatusSeeOther)
}