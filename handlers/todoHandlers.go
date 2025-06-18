package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"todo-app-4/model"
	"todo-app-4/services"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) HandleCreateTodoItem(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.CreateTodo(&todo)
	if err != nil {
		if err.Error() == "title is required" {
			http.Error(w, "title is required", http.StatusBadRequest)
		} else {
			log.Printf("Failed to create todo : %v \n", err)
			http.Error(w, "failed to create todo", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Printf("Failed to encode todo: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

func (h *TodoHandler) HandleGetAllTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		log.Printf("Failed to fetch todos: %v", err)
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		log.Printf("Failed to encode todos: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

func (h *TodoHandler) HandleGetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	todo, err := h.service.GetTodoById(id)
	if err != nil {
		if err.Error() == "todo not found" {
			http.Error(w, "todo not found", http.StatusNotFound)
		} else {
			log.Printf("Failed to fetch todo %v", err)
			http.Error(w, "failed to fetch todo", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Printf("Failed to encode: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TodoHandler) HandleDeleteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	err = h.service.DeleteTodoById(id)
	if err != nil {
		if err.Error() == "todo not found" {
			http.Error(w, "todo not found", http.StatusNotFound)
		} else {
			log.Printf("Failed to delete todo: %v", err)
			http.Error(w, "failed to delete to do", http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var updatedTodo model.Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	todo, err := h.service.UpdateTodoItem(id, &updatedTodo)
	if err != nil {
		if err.Error() == "todo not found" {
			http.Error(w, "todo not found", http.StatusNotFound)
		} else {
			log.Printf("Failed to update todo: %v", err)
			http.Error(w, "failed to update todo", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		log.Printf("Failed to encode %v", err)
		http.Error(w, "Failed to encode", http.StatusInternalServerError)
	}

}
