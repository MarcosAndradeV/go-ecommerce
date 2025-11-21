package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/MarcosAndradeV/go-ecommerce/internal/database"
	"github.com/MarcosAndradeV/go-ecommerce/internal/handlers"
	"github.com/MarcosAndradeV/go-ecommerce/internal/repository"
	"github.com/MarcosAndradeV/go-ecommerce/internal/routes"
	"github.com/MarcosAndradeV/go-ecommerce/internal/service"
)

func main() {
	// 1. Carregar variáveis
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando vars de ambiente")
	}

	port := os.Getenv("PORT")
	if port == "" { port = "8080" }
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" { log.Fatal("ERRO: MONGO_URI não definido") }

	// 2. Banco de Dados (Com Timeout para não travar)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	store := database.NewMongoStore(os.Getenv("DB_NAME"))
	if err := store.Connect(ctx, mongoURI); err != nil {
		log.Fatalf("FATAL: Falha ao conectar no MongoDB: %v", err)
	}
	defer func() {
		if err := store.Disconnect(context.Background()); err != nil {
			log.Printf("Erro ao desconectar: %v", err)
		}
	}()

	log.Println("Conectado ao MongoDB com sucesso!")

	// 3. Injeção de Dependências (Wiring)
	// Repositórios
	userRepo := repository.NewUserRepository(store.DB)
	storeRepo := repository.NewStoreRepository(store.DB)

	// Serviços (Aqui que o erro de nil poderia acontecer se userRepo fosse nil)
	authService := service.NewAuthService(userRepo)
	paymentService := service.NewPaymentService()
	storeService := service.NewStoreService(storeRepo, paymentService)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	storeHandler := handlers.NewStoreHandler(storeService)

	// 4. Rotas (Passamos authService também para o Middleware)
	r := routes.NewRouter(authHandler, storeHandler, authService)

	// 5. Servidor
	serverAddr := ":" + port
	log.Printf("Servidor rodando em http://localhost%s", serverAddr)

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
