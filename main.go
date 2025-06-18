package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
	"todo-app-4/handlers"
	"todo-app-4/repository"
	"todo-app-4/routes"
	"todo-app-4/services"
)

func main() {
	connStr := "postgres://postgres:password@localhost:5431/todo_app?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL!")
	todoRepo := repository.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	r := routes.SetRouter(todoHandler)
	handler := cors.Default().Handler(r)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}

}
