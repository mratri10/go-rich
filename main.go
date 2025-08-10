package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mratri10/go-rich/config"
	"github.com/mratri10/go-rich/handlers"
	"github.com/mratri10/go-rich/middleware"
	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://167.99.76.27:8080/"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	config.ConnectDB()

	r := chi.NewRouter()

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	// Protected routes
	r.Group(func(protected chi.Router) {
		protected.Use(middleware.AuthMiddleware)

		protected.Get("/profile", handlers.ProfileHandler)
		protected.Get("/user/all", handlers.UserAllHandler)
		protected.Post("/user/add", handlers.CreateUserHandler)
		protected.Post("/user/update", handlers.UpdateUserHandler)
		protected.Delete("/user/delete/{id}", handlers.DeleteUser)

		protected.Get("/role", handlers.GetRole)
		protected.Post("/role", handlers.AddRole)
		protected.Put("/role/{id}", handlers.UpdateRole)
		protected.Delete("/role/{id}", handlers.DeleteRole)

		protected.Get("/menu", handlers.GetMenu)
		protected.Post("/menu", handlers.AddMenu)
		protected.Put("/menu/{id}", handlers.UpdateMenu)
		protected.Delete("/menu/{id}", handlers.DeleteMenu)

		protected.Post("/menu/role", handlers.AddRoleMenu)
		protected.Post("/role/user", handlers.AddUserRole)
		protected.Delete("/role/user", handlers.DeleteUserRole)
		protected.Delete("/menu/role", handlers.DeleteRoleMenu)
	})

	fmt.Println("Server running on : 8080")
	http.ListenAndServe(":8080", c.Handler(r))
}
