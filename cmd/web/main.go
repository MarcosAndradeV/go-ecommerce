package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MarcosAndradeV/go-ecommerce/internal/database"
	"github.com/MarcosAndradeV/go-ecommerce/internal/handlers"
	"github.com/MarcosAndradeV/go-ecommerce/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbstore := database.NewMongoStore(env["DB_NAME"])

	if err := dbstore.Connect(ctx, env["MONGO_URI"]); err != nil {
		log.Println("Error: Não foi possivel connectar ao mongodb", err)
	}
	log.Println("Info: Connectado ao mongodb")

	r := chi.NewRouter()
	h := handlers.NewHandler(dbstore, ctx);
	routes.SetupRoutes(r, h)

	port := env["PORT"]
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, r)
	dbstore.Disconnect(ctx)
}
