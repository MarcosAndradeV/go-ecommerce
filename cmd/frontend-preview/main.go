package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	// Importe seus models reais para garantir que o template funcione
	// Troque "github.com/seu-usuario/go-ecommerce" pelo nome que está no seu go.mod
	"github.com/<user_name>/go-ecommerce/internal/models" 
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// 1. Servir CSS e Imagens (Essencial para o Tailwind funcionar)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Rota Home (Lista de Produtos Fake)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// DADOS MOCKADOS (Falsos)
		products := []models.Product{
			{
				ID:          primitive.NewObjectID(),
				Name:        "iPhone 15 Pro (Mock)",
				Description: "Titânio, chip A17 Pro, botão de ação.",
				ImageURL:    "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-15-pro-titanium-blue-select?wid=512&hei=512&fmt=jpeg&qlt=90&.v=1692891194305",
				Price:       999900, // R$ 9.999,00
				Stock:       10,
			},
			{
				ID:          primitive.NewObjectID(),
				Name:        "Notebook Gamer (Mock)",
				Description: "RTX 4060, 16GB RAM, SSD 1TB.",
				ImageURL:    "https://m.media-amazon.com/images/I/61Q-2n7A+rL._AC_SL1000_.jpg",
				Price:       540050, // R$ 5.400,50
				Stock:       0,      // Teste de Esgotado
			},
			{
				ID:          primitive.NewObjectID(),
				Name:        "Monitor UltraWide (Mock)",
				Description: "34 polegadas, 144hz, IPS.",
				ImageURL:    "https://m.media-amazon.com/images/I/71sxlhYhKWL._AC_SL1500_.jpg",
				Price:       250000, // R$ 2.500,00
				Stock:       5,
			},
		}
		renderTemplate(w, "index.html", products)
	})

	// 3. Rota Detalhe (Produto Único Fake)
	http.HandleFunc("/product/", func(w http.ResponseWriter, r *http.Request) {
		product := models.Product{
			ID:          primitive.NewObjectID(),
			Name:        "iPhone 15 Pro (Visualização Detalhe)",
			Description: "Esta é uma descrição longa simulada para testar a quebra de linha e o layout da página de detalhes. O acabamento em titânio é leve e robusto.",
			ImageURL:    "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-15-pro-titanium-blue-select?wid=512&hei=512&fmt=jpeg&qlt=90&.v=1692891194305",
			Price:       999900,
			Stock:       10,
		}
		renderTemplate(w, "product.html", product)
	})

	// 4. Rota Checkout (Simulação)
	http.HandleFunc("/checkout", func(w http.ResponseWriter, r *http.Request) {
		// Simula um produto sendo comprado
		product := models.Product{
			ID:       primitive.NewObjectID(),
			Name:     "iPhone 15 Pro",
			ImageURL: "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-15-pro-titanium-blue-select?wid=512&hei=512&fmt=jpeg&qlt=90&.v=1692891194305",
			Price:    999900,
		}
		renderTemplate(w, "checkout.html", product)
	})
	
	// 5. Rota Admin (Visualização)
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "admin.html", nil)
	})

	fmt.Println("🎨 Servidor de Frontend (MOCK) rodando em http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// Função Helper local para não depender dos handlers oficiais
func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	// Ajuste os caminhos se necessário, assumindo que roda da raiz do projeto
	t, err := template.ParseFiles("templates/layouts/base.html", "templates/"+tmplName)
	if err != nil {
		http.Error(w, "Erro Template: "+err.Error(), 500)
		return
	}
	t.ExecuteTemplate(w, "base", data)
}