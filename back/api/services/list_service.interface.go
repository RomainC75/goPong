package services

import (
	ListRequest "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	ListModel "github.com/saegus/test-technique-romain-chenard/data/models"
)

type ListServiceInterface interface {
	CreateList (list ListRequest.CreateListRequest, userId string) (ListModel.List, error)
	GetList (listId string) (ListModel.List, error)
	GetLists (userId string) ([]ListModel.List, error)
	DeleteList (userId string, listId string) (ListModel.List, error)
	UpdateList (userId string, list ListModel.List) (ListModel.List, error)
	IsUserTheOwnerOfTHeList(userId string, listId string) (bool, error)
}
