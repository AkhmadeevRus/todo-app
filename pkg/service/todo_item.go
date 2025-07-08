package service

import (
	"github.com/akhmadeevrus/todo-app"
	"github.com/akhmadeevrus/todo-app/pkg/repository"
)

type TodoItemServise struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemServise {
	return &TodoItemServise{repo: repo, listRepo: listRepo}
}

func (s *TodoItemServise) Create(userId, id int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, id)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(id, item)
}
