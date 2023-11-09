package repositories

import (
	"errors"
	"fmt"

	models "github.com/saegus/test-technique-romain-chenard/internal/modules/list/models"
	database "github.com/saegus/test-technique-romain-chenard/pkg/database"
	"gorm.io/gorm"
)

type ListRepository struct {
	DB *gorm.DB
}

func New() *ListRepository{
	return  &ListRepository{
		DB: database.Connection(),
	}
}

func (ListRepository *ListRepository) CreateList(list models.List) (models.List, error){
	var newList models.List
	result := ListRepository.DB.Create(&list).Scan(&newList)
	if result.RowsAffected == 0 {
		return models.List{}, errors.New("error trying to creat a new user")
	}
	return newList, nil
}

func (ListRepository *ListRepository) GetLists(userId string) ([]models.List, error){
	var foundLists []models.List
	result := ListRepository.DB.Where("user_id = ?", userId).Find(&foundLists)
	if result.RowsAffected == 0 {
		return []models.List{}, errors.New("error trying to get the lists")
	}
	return foundLists, nil
}

func (ListRepository *ListRepository) GetListById(listId string) (models.List, error){
	fmt.Println("==> ", listId)
	var foundList models.List
	result := ListRepository.DB.Where("id = ?", listId).First(&foundList)	
	if result.RowsAffected == 0 {
		return models.List{}, errors.New(fmt.Sprintf("error trying to get the list : %s ", listId))
	}
	return foundList, nil
}

func (ListRepository *ListRepository) DeleteList(userId string, listId string) (models.List, error){
	var deletedList models.List

	if err := ListRepository.DB.Where("user_id = ?", userId).Where("id = ?", listId).First(&deletedList).Error; err != nil {
        if gorm.ErrRecordNotFound == err {
            return models.List{}, errors.New("List item not found")
        }
        return models.List{}, err
    }

	if err := ListRepository.DB.Delete(&deletedList).Error; err != nil {
        return models.List{}, err
    }

	return deletedList, nil
}

func (ListRepository *ListRepository) UpdateList(userId string, list models.List) (models.List, error){
	var updatedList models.List

	if err := ListRepository.DB.First(&updatedList, list.ID).Error; err != nil {
        if gorm.ErrRecordNotFound == err {
            return models.List{}, errors.New("List item not found")
        }
        return models.List{}, err
    }

    updatedList.Name = list.Name
    
    if err := ListRepository.DB.Save(&updatedList).Error; err != nil {
        return models.List{}, err
    }

	return updatedList, nil
}

