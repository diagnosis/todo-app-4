package services

import (
	"errors"
	"todo-app-4/model"
	"todo-app-4/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo}
}

func (s *TodoService) CreateTodo(todo *model.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}

	return s.repo.InsertInto(todo)
}

func (s *TodoService) GetAllTodos() ([]model.Todo, error) {
	return s.repo.GetAllRows()
}
func (s *TodoService) GetTodoById(id int) (*model.Todo, error) {
	return s.repo.GetTodoByID(id)
}

func (s *TodoService) DeleteTodoById(id int) error {
	return s.repo.DeleteTodoByID(id)
}

func (s *TodoService) UpdateTodoItem(id int, updatedTodo *model.Todo) (*model.Todo, error) {
	if updatedTodo.Title == "" {
		return nil, errors.New("title is required")
	}
	todo := &model.Todo{ID: id}
	err := s.repo.UpdateTodo(todo, updatedTodo)
	if err != nil {
		return nil, err
	}
	return s.repo.GetTodoByID(id)
}
