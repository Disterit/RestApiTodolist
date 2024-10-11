package service

import (
	todo "ResTApiTodolist"
	"ResTApiTodolist/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
