package repository

import (
	"github.com/MuShaf-NMS/go_gin/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todo)
	GetAllTodos(userID uuid.UUID) []models.Todo
	GetOneTodo(uuid uuid.UUID) (models.Todo, error)
	UpdateTodo(todo *models.Todo, uuid uuid.UUID)
	DeleteTodo(uuid uuid.UUID)
}

type todoRepository struct {
	connection *gorm.DB
}

func (tr *todoRepository) CreateTodo(todo *models.Todo) {
	tr.connection.Create(todo)
}

func (tr *todoRepository) GetAllTodos(userID uuid.UUID) []models.Todo {
	var todos []models.Todo
	tr.connection.Where(&models.Todo{UserID: userID}).Find(&todos)
	return todos
}

func (tr *todoRepository) GetOneTodo(uuid uuid.UUID) (models.Todo, error) {
	var todo models.Todo
	err := tr.connection.Where(models.Todo{UUID: uuid}).First(&todo).Error
	return todo, err
}

func (tr *todoRepository) UpdateTodo(todo *models.Todo, uuid uuid.UUID) {
	tr.connection.Model(models.Todo{}).Where("uuid = ?", uuid).Updates(todo)
}

func (tr *todoRepository) DeleteTodo(uuid uuid.UUID) {
	tr.connection.Delete(&models.Todo{}, "uuid = ?", uuid.String())
}

func NewTodoRepository(dbConnection *gorm.DB) TodoRepository {
	return &todoRepository{
		connection: dbConnection,
	}
}
