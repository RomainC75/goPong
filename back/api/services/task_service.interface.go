package services

import (
	"github.com/saegus/test-technique-romain-chenard/internal/modules/task/models"
	TaskModel "github.com/saegus/test-technique-romain-chenard/internal/modules/task/models"
	TaskRequest "github.com/saegus/test-technique-romain-chenard/internal/modules/task/requests"
)

type TaskServiceInterface interface {
	CreateTask (task TaskRequest.CreateTaskRequest, listId string) (TaskModel.Task, error)
	GetTasks (userId string) []TaskModel.Task
	GetTask (taskId string) (models.Task, error)
	ToggleTaskIsDone (taskId string) (models.Task, error)
	UpdateTask(task models.Task) (models.Task, error)
	Delete(taskId string) (models.Task, error)
	DeleteTasksListId (listId string) ([]models.Task, error)
}
