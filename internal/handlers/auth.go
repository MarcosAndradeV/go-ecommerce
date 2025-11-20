package handlers

import (
	"net/http"
	"time"

	"github.com/MarcosAndradeV/go-ecommerce/internal/service"
)

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "register.html", nil)
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Chama o service para criar o usuário no Mongo
	err := service.RegisterCustomer(name, email, password)
	if err != nil {
		// Se der erro (ex: email duplicado), volta pro form
		http.Redirect(w, r, "/register?error=true", http.StatusSeeOther)
		return
	}

	// Sucesso! Vai para o login
	http.Redirect(w, r, "/login?success=created", http.StatusSeeOther)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	// --- CAMINHO A: É O ADMIN? (Fixo) ---
	if email == "admin" && password == "admin123" {
		// Cria cookie específico de Admin
		http.SetCookie(w, &http.Cookie{
			Name:    "sessao_admin",
			Value:   "true",
			Path:    "/",
			Expires: time.Now().Add(24 * time.Hour),
		})
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	// --- CAMINHO B: É UM CLIENTE? (Banco de Dados) ---
	user, err := service.AuthenticateCustomer(email, password)
	if err == nil {
		// Sucesso! Cria cookie de Cliente com o email dele
		http.SetCookie(w, &http.Cookie{
			Name:    "sessao_loja",
			Value:   user.Email, // Guardamos o email pra saber quem é
			Path:    "/",
			Expires: time.Now().Add(24 * time.Hour),
		})
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// --- CAMINHO C: SENHA ERRADA ---
	http.Redirect(w, r, "/login?error=invalid", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Destroi o cookie de Admin
	http.SetCookie(w, &http.Cookie{
		Name:   "sessao_admin",
		MaxAge: -1,
	})

	// Destroi o cookie de Cliente
	http.SetCookie(w, &http.Cookie{
		Name:   "sessao_loja",
		MaxAge: -1,
	})

	// Manda pra Home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
