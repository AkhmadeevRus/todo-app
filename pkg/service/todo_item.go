package service

import (
	"github.com/akhmadeevrus/todo-app"
	"github.com/akhmadeevrus/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, id int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, id)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(id, item)
}

func (s *TodoItemService) GetAll(userId, id int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, id)
}

func (s *TodoItemService) GetById(userId, id int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, id)
}

func (s *TodoItemService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}

func (s *TodoItemService) Update(userId, id int, input todo.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}
