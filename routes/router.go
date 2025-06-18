package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"todo-app-4/handlers"
)

func SetRouter(handler *handlers.TodoHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Todo API"))
	})
	r.Post("/todos", handler.HandleCreateTodoItem)
	r.Get("/todos", handler.HandleGetAllTodo)
	r.Get("/todos/{id}", handler.HandleGetTodoByID)
	r.Delete("/todos/{id}", handler.HandleDeleteByID)
	r.Put("/todos/{id}", handler.HandleUpdate)
	return r
}
