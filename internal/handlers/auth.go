package handlers

import (
	"net/http"
	"time"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, r, "login.html", nil)
}

func PerformLogin(w http.ResponseWriter, r *http.Request) {
	senha := r.FormValue("password")

	// Senha Fixa (Em produção use variáveis de ambiente)
	if senha == "admin123" {
		http.SetCookie(w, &http.Cookie{
			Name:    "sessao_admin",
			Value:   "logado",
			Expires: time.Now().Add(24 * time.Hour),
			Path:    "/",
		})
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		// Redireciona com erro (simples) ou renderiza com msg
		http.Redirect(w, r, "/login?erro=senha", http.StatusSeeOther)
	}
}

func PerformLogout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "sessao_admin",
		MaxAge: -1, // Mata o cookie
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}