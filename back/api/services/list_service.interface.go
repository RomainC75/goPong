package services

import (
	ListModel "github.com/saegus/test-technique-romain-chenard/internal/modules/list/models"
	ListRequest "github.com/saegus/test-technique-romain-chenard/internal/modules/list/requests"
)

type ListServiceInterface interface {
	CreateList (list ListRequest.CreateListRequest, userId string) (ListModel.List, error)
	GetList (listId string) (ListModel.List, error)
	GetLists (userId string) ([]ListModel.List, error)
	DeleteList (userId string, listId string) (ListModel.List, error)
	UpdateList (userId string, list ListModel.List) (ListModel.List, error)
	IsUserTheOwnerOfTHeList(userId string, listId string) (bool, error)
}
