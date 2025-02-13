package app

import (
	"context"
	"todo-app/internal/models"
	"todo-app/internal/service"
)

type App struct {
	todoService service.TodoService
	ctx         context.Context
}

func NewApp(todoService service.TodoService) *App {
	return &App{todoService: todoService}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AddTodo(title string, description string, dueDateStr string, priority string) error {
	err := a.todoService.AddTodo(title, description, dueDateStr, priority)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetTodos() ([]models.Todo, error) {
	todos, err := a.todoService.GetTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (a *App) ToggleTodo(id int) error {
	err := a.todoService.ToggleTodo(id)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) DeleteTodo(id int) error {
	_, err := a.todoService.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}

// func StartApp(todoService service.TodoService) error {
// 	app := NewApp(todoService)

// 	err := wails.Run(&wails.AppConfig{
// 		Width:  1024,
// 		Height: 768,
// 		Title:  "Todo App",
// 		JS:     []string{"frontend/dist/app.js"},
// 		CSS:    []string{"frontend/dist/app.css"},
// 		Colour: "#131313",
// 		Bind: []interface{}{
// 			app,
// 		},
// 	})

// 	if err != nil {
// 		return fmt.Errorf("failed to start Wails app: %v", err)
// 	}

// 	return nil
// }
