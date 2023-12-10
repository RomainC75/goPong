package services

import (
	Requests "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	Models "github.com/saegus/test-technique-romain-chenard/data/models"
)

type TaskServiceInterface interface {
	CreateTask (task Requests.CreateTaskRequest, listId string) (Models.Task, error)
	GetTasks (userId string) []Models.Task
	GetTask (taskId string) (Models.Task, error)
	ToggleTaskIsDone (taskId string) (Models.Task, error)
	UpdateTask(task Models.Task) (Models.Task, error)
	Delete(taskId string) (Models.Task, error)
	DeleteTasksListId (listId string) ([]Models.Task, error)
}
