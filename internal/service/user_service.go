package service

import (
	"errors"
	"time"

	"github.com/MarcosAndradeV/go-ecommerce/internal/models"
	"github.com/MarcosAndradeV/go-ecommerce/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
)

// Registra um cliente novo
func RegisterCustomer(db *mongo.Database, name, email, password string) error {
	// 1. Verifica se já existe
	existing, _ := repository.GetUserByEmail(db, email)
	if existing != nil {
		return errors.New("este e-mail já está cadastrado")
	}

	// 2. Hash da senha (Segurança)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 3. Cria o objeto
	user := models.User{
		ID:           primitive.NewObjectID(),
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPass),
		IsAdmin:      false, // Clientes nunca nascem admin
		CreatedAt:    time.Now(),
	}

	// 4. Salva
	return repository.CreateUser(db, user)
}

// Autentica o usuário (Login)
func AuthenticateCustomer(db *mongo.Database, email, password string) (*models.User, error) {
	// 1. Busca usuário
	user, err := repository.GetUserByEmail(db, email)
	if err != nil {
		return nil, errors.New("usuário ou senha inválidos")
	}

	// 2. Compara Hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("usuário ou senha inválidos")
	}

	return user, nil
}

// Dados para o Dashboard
func GetCustomerDashboard(db *mongo.Database, email string) (*models.User, []models.Order, error) {
	user, err := repository.GetUserByEmail(db, email)
	if err != nil {
		return nil, nil, err
	}

	orders, err := repository.GetOrdersByEmail(db, email)
	if err != nil {
		// Se der erro ao buscar pedidos, retorna lista vazia, mas não trava o user
		orders = []models.Order{}
	}

	return user, orders, nil
}
