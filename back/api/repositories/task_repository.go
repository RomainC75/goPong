package repositories

import (
	"errors"

	database "github.com/saegus/test-technique-romain-chenard/data/database"
	models "github.com/saegus/test-technique-romain-chenard/data/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepo() *TaskRepository{
	return  &TaskRepository{
		DB: database.Connection(),
	}
}

func (taskRepository *TaskRepository) CreateTask(task models.Task) (models.Task, error){
	var newTask models.Task
	result := taskRepository.DB.Create(&task).Scan(&newTask)
	if result.RowsAffected == 0 {
		return models.Task{}, errors.New("error trying to create a new task")
	}
	return newTask, nil
}

func (taskRepository *TaskRepository) GetTasks(listId string) []models.Task{
	var foundTasks []models.Task
	taskRepository.DB.Where("list_id = ? ", listId).Find(&foundTasks)
	return foundTasks
}

func (taskRepository *TaskRepository) GetTaskById(taskId string) (models.Task, error){
	var foundTask models.Task
	if err := taskRepository.DB.Where("id = ?", taskId).First(&foundTask).Error; err != nil {
		return models.Task{}, err
	}
	return foundTask, nil
}

func (TaskRepository *TaskRepository) ToggleTaskIsDoneById (taskId string) (models.Task, error) {
	var foundTask models.Task
	if err := TaskRepository.DB.Where("id = ?", taskId).First(&foundTask).Error; err != nil{
		return models.Task{}, errors.New("error trying to get the task")
	}
	foundTask.IsDone = !foundTask.IsDone
	if err := TaskRepository.DB.Save(&foundTask).Error; err != nil {
		return models.Task{}, errors.New("error trying to update the task")
	}
	return foundTask, nil
}

func (TaskRepository *TaskRepository) UpdateTask (task models.Task) (models.Task, error) {
	var foundTask models.Task
	if err := TaskRepository.DB.Where("id = ?", task.ID.String()).First(&foundTask).Error ; err != nil {
		return models.Task{}, errors.New("error trying to get the task")
	}
	foundTask.Name = task.Name
	foundTask.Description = task.Description
	foundTask.IsDone = task.IsDone
	foundTask.DeadLine = task.DeadLine

	if err := TaskRepository.DB.Save(&foundTask).Error ; err != nil {
		return models.Task{}, errors.New("error trying to update the task")
	}

	return foundTask, nil
}

func (TaskRepository *TaskRepository) DeleteTaskById (taskId string) (models.Task, error) {
	var foundTask models.Task
	if err := TaskRepository.DB.Where("id = ?", taskId).First(&foundTask).Error; err != nil{
		return models.Task{}, err
	}
	if err := TaskRepository.DB.Delete(&foundTask).Error; err != nil{
		return models.Task{}, err
	}
	return foundTask, nil
}

func (TaskRepository *TaskRepository) DeleteTasksByListId (listId string) ([]models.Task, error){
	var deletedTasks []models.Task
	if err := TaskRepository.DB.Where("list_id = ?", listId).Delete(deletedTasks).Error; err != nil {
		return []models.Task{}, err
	}

	return deletedTasks, nil
	
}