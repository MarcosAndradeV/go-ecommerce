package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// renderTemplate junta o layout base com o template específico
func RenderTemplate(w http.ResponseWriter, tmplName string, data any) {
	// Caminhos dos arquivos
	layoutPath := filepath.Join("templates", "layouts", "base.html")
	templatePath := filepath.Join("templates", tmplName)

	// Parse dos arquivos
	tmpl, err := template.ParseFiles(layoutPath, templatePath)
	if err != nil {
		http.Error(w, "Erro interno no servidor (Template): "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Executa o template "base" (definido no base.html), passando os dados
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Erro ao renderizar página: "+err.Error(), http.StatusInternalServerError)
	}
}
