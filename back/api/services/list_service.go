package services

import (
	"github.com/google/uuid"
	ListModel "github.com/saegus/test-technique-romain-chenard/internal/modules/list/models"
	ListRepository "github.com/saegus/test-technique-romain-chenard/internal/modules/list/repositories"
	ListRequest "github.com/saegus/test-technique-romain-chenard/internal/modules/list/requests"
)

type ListService struct {
	listRepository ListRepository.ListRepositoryInterface
}

func New() *ListService{
	return &ListService{
		listRepository: ListRepository.New(),
	}
}

func (listService *ListService) CreateList (list ListRequest.CreateListRequest, userId string) (ListModel.List, error){
	var newList ListModel.List
	userUuid, _ := uuid.Parse(userId)
	newList.Name = list.Name
	newList.UserId = userUuid
	newList, err := listService.listRepository.CreateList(newList)
	if err != nil{
		return ListModel.List{}, err
	}
	return newList, nil
}

func (listService *ListService) GetLists (userId string) ([]ListModel.List, error){
	lists, err := listService.listRepository.GetLists(userId)
	if err != nil {
		return []ListModel.List{}, err
	}
	return lists, nil
}

func (listService *ListService) GetList (listId string) (ListModel.List, error){
	list, err := listService.listRepository.GetListById(listId)
	if err != nil {
		return ListModel.List{}, err
	}
	return list, nil
}

func (listService *ListService) DeleteList (userId string, listId string) (ListModel.List, error){
	list, err := listService.listRepository.DeleteList(userId, listId)
	if err != nil {
		return ListModel.List{}, err
	}
	return list, nil
}

func (listService *ListService) UpdateList (userId string, list ListModel.List) (ListModel.List, error){
	list, err := listService.listRepository.UpdateList(userId, list)
	if err != nil {
		return ListModel.List{}, err
	}
	return list, nil
}

func (listService *ListService) IsUserTheOwnerOfTHeList (userId string, listId string) (bool, error){
	foundList, err := listService.listRepository.GetListById(listId)
	if  err != nil{
		return false, err
	}
	
	if userId != foundList.UserId.String(){
		return false, nil
	}
	return true, nil
}
