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

	r.Get("/menu", middleware.AuthMiddleware(handlers.GetMenu))
	r.Post("/menu", middleware.AuthMiddleware(handlers.AddMenu))
	r.Put("/menu/{id}", middleware.AuthMiddleware(handlers.UpdateMenu))
	r.Delete("/menu/{id}", middleware.AuthMiddleware(handlers.DeleteMenu))

	r.Post("/menu/role", middleware.AuthMiddleware(handlers.AddRoleMenu))
	r.Post("/role/user", middleware.AuthMiddleware(handlers.AddUserRole))
	r.Delete("/role/user", middleware.AuthMiddleware(handlers.DeleteUserRole))
	r.Delete("/menu/role", middleware.AuthMiddleware(handlers.DeleteRoleMenu))

	fmt.Println("Server running on : 8080")
	http.ListenAndServe(":8080", r)
}
