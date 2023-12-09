package repositories

import (
	models "github.com/saegus/test-technique-romain-chenard/internal/modules/list/models"
)

type ListRepositoryInterface interface {
	CreateList(list models.List) (models.List, error)
	GetLists(userId string) ([]models.List, error)
	GetListById(listId string) (models.List, error)
	DeleteList(userId string, listId string) (models.List, error)
	UpdateList(userId string, list models.List) (models.List, error)
}
