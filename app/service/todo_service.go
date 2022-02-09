package service

import (
	"github.com/MuShaf-NMS/go_gin/app/dto"
	"github.com/MuShaf-NMS/go_gin/app/models"
	"github.com/MuShaf-NMS/go_gin/app/repository"
	"github.com/google/uuid"
)

type TodoService interface {
	CreateTodo(todo dto.CreateTodo, userId uuid.UUID) models.Todo
	GetOneTodo(uuid uuid.UUID) (models.Todo, error)
	GetAllTodos(userId uuid.UUID) []models.Todo
	UpdateTodo(todo dto.UpdateTodo, uuid uuid.UUID)
	DeleteTodo(uuid uuid.UUID)
}

type todoService struct {
	repository repository.TodoRepository
}

func (ts *todoService) CreateTodo(todo dto.CreateTodo, userID uuid.UUID) models.Todo {
	todoCreate := models.Todo{
		Name:        todo.Name,
		Category:    todo.Category,
		Description: todo.Description,
		UserID:      userID,
	}
	ts.repository.CreateTodo(&todoCreate)
	return todoCreate
}

func (ts *todoService) GetOneTodo(uuid uuid.UUID) (models.Todo, error) {
	todo, err := ts.repository.GetOneTodo(uuid)
	return todo, err
}

func (ts *todoService) GetAllTodos(userId uuid.UUID) []models.Todo {
	todos := ts.repository.GetAllTodos(userId)
	return todos
}

func (ts *todoService) UpdateTodo(todo dto.UpdateTodo, uuid uuid.UUID) {
	todoUpdate := models.Todo{
		Name:        todo.Name,
		Category:    todo.Category,
		Description: todo.Description,
	}
	ts.repository.UpdateTodo(&todoUpdate, uuid)
}

func (ts *todoService) DeleteTodo(uuid uuid.UUID) {
	ts.repository.DeleteTodo(uuid)
}

func NewTodoService(repository repository.TodoRepository) TodoService {
	return &todoService{
		repository: repository,
	}
}
