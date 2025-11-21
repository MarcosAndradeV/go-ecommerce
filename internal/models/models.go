package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	IsAdmin      bool               `bson:"is_admin"`
	CreatedAt    time.Time          `bson:"created_at"`	
	Cart 		[]OrderItem 		`bson:"cart,omitempty"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	ImageURL    string             `bson:"image_url"`

	Price int64 `bson:"price"`

	Stock int `bson:"stock"`
	Sizes []string `bson:"sizes"` // <--- Generic Size/Attribute

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (p Product) FormattedPrice() string {
	return fmt.Sprintf("R$ %.2f", float64(p.Price)/100)
}

func (p Product) PriceToFloat() float64 {
	return float64(p.Price) / 100.0
}

type OrderItem struct {
	ProductID   primitive.ObjectID `bson:"product_id"`
	ProductName string             `bson:"product_name"`

	Price int64 										`bson:"price"`
	Quantity int 										`bson:"quantity"`
	Size     string             `bson:"size"` // <--- Selected Size
	ImageURL    string             	`bson:"image_url"`
}

func (i OrderItem) TotalItem() string {
	total := i.Price * int64(i.Quantity)
	return fmt.Sprintf("R$ %.2f", float64(total)/100)
}

type Order struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`

	CustomerName    string `bson:"customer_name"`
	CustomerEmail   string `bson:"customer_email"`
	CustomerAddress string `bson:"customer_address,omitempty"`

	Items []OrderItem `bson:"items"`

	Total  int64  `bson:"total"`
	Status string `bson:"status"`

	CreatedAt time.Time `bson:"created_at"`
}

func (o Order) FormattedTotal() string {
	return fmt.Sprintf("R$ %.2f", float64(o.Total)/100)
}
