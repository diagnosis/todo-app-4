package model

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DueDate     time.Time `json:"due_date"`
	GroupName   *string   `json:"groupName,omitempty"`
	Description *string   `json:"description,omitempty"`
}

func NewTodo() *Todo {
	return &Todo{}
}

func (t *Todo) TimeLeft() string {
	if t.DueDate.IsZero() {
		return "No due date"
	}
	remaining := time.Until(t.DueDate)
	if remaining < 0 {
		return "Overdue"
	}
	return remaining.Round(time.Hour).String()
}
