package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MarcosAndradeV/go-ecommerce/internal/database"
	"github.com/MarcosAndradeV/go-ecommerce/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	product := models.Product{
		ID:          primitive.NewObjectID(),
		Name:        "iPhone 15 Pro (Mock)",
		Description: "Titânio, chip A17 Pro, botão de ação.",
		ImageURL:    "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-15-pro-titanium-blue-select?wid=512&hei=512&fmt=jpeg&qlt=90&.v=1692891194305",
		Price:       999900,
		Stock:       10,
	}
	if _, err := dbstore.DB.Collection("Product").InsertOne(ctx, product); err != nil {
		log.Println("Error:", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	port := env["PORT"]
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Servidor rodando em http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, r)
	dbstore.Disconnect(ctx)
}
