package main

import (
	"log"
	"net/http"
	"trueAPI/internal/handlers"
	"trueAPI/internal/repository"
	"trueAPI/internal/services"
)

func main() {
	repo := repository.NewUserRepository()
	userService := services.NewUserService(repo)
	userHandler := handlers.NewUserHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Query().Has("id") {
				userHandler.GetUserByID(w, r)
				return
			}
			userHandler.GetAllUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("trueAPI listening on http://localhost:8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}
