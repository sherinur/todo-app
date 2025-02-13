package service

import (
	"time"
	"todo-app/internal/dal"
	"todo-app/internal/models"
)

type TodoService interface {
	AddTodo(title string, description string, dueDateStr string, priority string) error
	GetTodos() ([]models.Todo, error)
	ToggleTodo(id int) error
	DeleteTodo(id int) (models.Todo, error)
}

type todoService struct {
	repo dal.TodoRepo
}

func NewTodoService(repo dal.TodoRepo) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) AddTodo(title string, description string, dueDateStr string, priority string) error {
	if len(title) == 0 {
		return ErrEmptyTitle
	}

	if len(description) == 0 {
		return ErrEmptyDescription
	}

	dueDate, err := time.Parse(time.RFC3339, dueDateStr)
	if err != nil {
		return err
	}

	if priority != "low" && priority != "medium" && priority != "high" {
		return ErrInvalidPriority
	}

	todo := models.Todo{
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		DueDate:     dueDate,
		Priority:    priority,
	}

	_, err = s.repo.AddTodo(todo)
	if err != nil {
		return err
	}

	return nil
}

func (s *todoService) GetTodos() ([]models.Todo, error) {
	return s.repo.GetTodos()
}

func (s *todoService) ToggleTodo(id int) error {
	return s.repo.ToggleTodo(id)
}

func (s *todoService) DeleteTodo(id int) (models.Todo, error) {
	return s.repo.DeleteTodo(id)
}
