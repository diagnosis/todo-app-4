package repository

import (
	"database/sql"
	"errors"
	"todo-app-4/model"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (todoRepo *TodoRepository) InsertInto(todo *model.Todo) error {
	query := `INSERT INTO todos (title, completed, due_date, groupName, description) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at, due_date, groupName, description`
	err := todoRepo.db.QueryRow(query, todo.Title, todo.Completed, todo.DueDate, todo.GroupName, todo.Description).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt, &todo.DueDate, &todo.GroupName, &todo.Description)
	if err != nil {
		return err
	}
	return nil
}

func (todoRepo *TodoRepository) GetAllRows() ([]model.Todo, error) {
	rows, err := todoRepo.db.Query(`SELECT id, title, completed, created_at, updated_at, due_date, groupName, description FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DueDate, &todo.GroupName, &todo.Description)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (todoRepo *TodoRepository) GetTodoByID(id int) (*model.Todo, error) {
	var todo model.Todo
	query := `SELECT id, title, completed, created_at, updated_at, due_date, groupName, description FROM todos WHERE id = $1`
	err := todoRepo.db.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DueDate, &todo.GroupName, &todo.Description)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("todo not found")
	}
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (todoRepo *TodoRepository) DeleteTodoByID(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	result, err := todoRepo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}

func (todoRepo *TodoRepository) UpdateTodo(todo *model.Todo, updatedTodo *model.Todo) error {
	query := `UPDATE todos SET title = $1, completed = $2, due_date = $3, groupName = $4, description = $5 WHERE id = $6`
	result, err := todoRepo.db.Exec(query, updatedTodo.Title, updatedTodo.Completed, updatedTodo.DueDate, updatedTodo.GroupName, updatedTodo.Description, todo.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}
