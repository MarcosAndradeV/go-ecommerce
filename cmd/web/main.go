package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MarcosAndradeV/go-ecommerce/internal/database"
	"github.com/MarcosAndradeV/go-ecommerce/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado")
	}

	db := database.NewMongoStore(env["DB_NAME"])

	if err := db.Connect(env["MONGO_URI"]); err != nil {
		log.Println("Error: Não foi possivel connectar ao mongodb")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.RenderTemplate(w, "../templates/index.html", nil)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, r)
}
