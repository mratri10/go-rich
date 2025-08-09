package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mratri10/go-rich/config"
	"github.com/mratri10/go-rich/handlers"
	"github.com/mratri10/go-rich/middleware"
)

func main() {
	config.ConnectDB()

	r := chi.NewRouter()

	r.Post("/register", handlers.RegisterHandler)
	r.Post("/login", handlers.LoginHandler)

	r.Get("/profile", middleware.AuthMiddleware(handlers.ProfileHandler))
	r.Get("/user/all", middleware.AuthMiddleware(handlers.UserAllHandler))
	r.Post("/user/add", middleware.AuthMiddleware(handlers.CreateUserHandler))
	r.Post("/user/update", middleware.AuthMiddleware(handlers.UpdateUserHandler))
	r.Delete("/user/delete/{id}", middleware.AuthMiddleware(handlers.DeleteUser))

	r.Get("/role", middleware.AuthMiddleware(handlers.GetRole))
	r.Post("/role", middleware.AuthMiddleware(handlers.AddRole))
	r.Put("/role/{id}", middleware.AuthMiddleware(handlers.UpdateRole))
	r.Delete("/role/{id}", middleware.AuthMiddleware(handlers.DeleteRole))

	fmt.Println("Server running on : 8080")
	http.ListenAndServe(":8080", r)
}
