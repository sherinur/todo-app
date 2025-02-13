package dal

import (
	"fmt"
	"todo-app/internal/models"

	"gorm.io/gorm"
)

type TodoRepo interface {
	AddTodo(todo models.Todo) (models.Todo, error)
	GetTodos() ([]models.Todo, error)
	ToggleTodo(id int) error
	DeleteTodo(id int) (models.Todo, error)
}

type todoRepo struct {
	db    *gorm.DB
	todos []models.Todo
	count uint
}

func NewTodoRepo() TodoRepo {
	return &todoRepo{}
}

func (r *todoRepo) AddTodo(todo models.Todo) (models.Todo, error) {
	r.todos = append(r.todos, todo)
	r.count++
	return todo, nil
}

func (r *todoRepo) GetTodos() ([]models.Todo, error) {
	return r.todos, nil
}

func (r *todoRepo) ToggleTodo(id int) error {
	for _, todo := range r.todos {
		if todo.ID == id {
			r.todos[id].Done = !r.todos[id].Done
			return nil
		}
	}

	return fmt.Errorf("todo with id %d not found", id)
}

func (r *todoRepo) DeleteTodo(id int) (models.Todo, error) {
	for i, todo := range r.todos {
		if todo.ID == id {
			deletedTodo := r.todos[i]
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			r.count--
			return deletedTodo, nil
		}
	}

	return models.Todo{}, fmt.Errorf("todo with id %d not found", id)
}
