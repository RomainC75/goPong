package responses

import (
	"time"

	"github.com/google/uuid"

	TaskModel "github.com/saegus/test-technique-romain-chenard/data/models"
)


type TaskResponse struct {
	ID       uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	IsDone bool `json:"isDone"`
	Description string `json:"description"`
	DeadLine time.Time `json:"deadLine"`
}


func ToTaskResponse (task TaskModel.Task)TaskResponse{
	return TaskResponse{
		ID: task.ID,
		Name: task.Name,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeadLine: task.DeadLine,
		Description: task.Description,
		IsDone: task.IsDone,
	}
}

func ToTaskArrayResponse (taskArray []TaskModel.Task) []TaskResponse{
	taskResponse := make([]TaskResponse, 0)
	for _, t := range(taskArray){
		taskResponse = append(taskResponse, ToTaskResponse(t))
	}
	return taskResponse
}