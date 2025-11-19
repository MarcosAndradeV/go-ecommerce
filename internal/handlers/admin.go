package handlers

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarcosAndradeV/go-ecommerce/internal/database"
	"github.com/MarcosAndradeV/go-ecommerce/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Dashboard Admin
func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	// Verifica Login
	if !CheckAuth(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	coll := database.GetCollection("products")
	cursor, _ := coll.Find(context.TODO(), bson.M{})
	
	var products []models.Product
	cursor.All(context.TODO(), &products)

	renderTemplate(w, r, "admin.html", products)
}

// Criar Produto
func AdminCreateProduct(w http.ResponseWriter, r *http.Request) {
	if !CheckAuth(r) {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	// Conversão de R$ string para int64 centavos
	priceStr := r.FormValue("price") // Ex: "10.50" ou "10,50"
	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	priceFloat, _ := strconv.ParseFloat(priceStr, 64)
	priceInt := int64(priceFloat * 100)

	stock, _ := strconv.Atoi(r.FormValue("stock"))

	product := models.Product{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		ImageURL:    r.FormValue("image_url"),
		Price:       priceInt,
		Stock:       stock,
	}

	coll := database.GetCollection("products")
	coll.InsertOne(context.TODO(), product)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}