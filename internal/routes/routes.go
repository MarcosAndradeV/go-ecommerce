package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *chi.Mux, db *mongo.Database) {

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	// // Loja (público)
	// r.Get("/", h.HomeHandler)
	// r.Get("/product{id}", h.ProductDetailHandler)
	// //Carrinho
	// r.Get("/checkout", h.CheckoutPageHandler)
	// r.Post("/purchase", h.PurchaseHandler)
	// r.Get("/sucess", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.RenderTemplate(w, r, "sucess.html", nil)
	// })
	// //ADMIM
	// r.Get("/login", h.LoginPage)
	// r.Get("/do-login", h.PerformLogin)
	// r.Get("/logout", h.PerformLogout)

	// r.Route("/admin", func(r chi.Router){
	// 	r.Get("/", h.AdminDashboard)
	// 	r.Get("/create", h.AdminCreateProduct)
	// })
}
