package services

import (
	"github.com/google/uuid"

	TaskRequest "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	TaskRepository "github.com/saegus/test-technique-romain-chenard/api/repositories"
	Models "github.com/saegus/test-technique-romain-chenard/data/models"
)

type TaskService struct {
	taskRepository TaskRepository.TaskRepositoryInterface
}

func NewTaskSrv() *TaskService{
	return &TaskService{
		taskRepository: TaskRepository.NewTaskRepo(),
	}
}

func (taskService *TaskService) CreateTask (task TaskRequest.CreateTaskRequest, listId string) (Models.Task, error){
	var newTask Models.Task
	listUuid := uuid.MustParse(listId)

	newTask.Name = task.Name
	newTask.ListId = listUuid
	newTask.DeadLine = task.DeadLine
	newTask.Description = task.Description
	newTask.IsDone = false
	
	newTask, err := taskService.taskRepository.CreateTask(newTask)

	if err != nil{
		return Models.Task{}, err
	}
	return newTask, nil
}


func (taskService *TaskService) GetTasks (listId string) []Models.Task{
	return taskService.taskRepository.GetTasks(listId)
}

func (taskService *TaskService) GetTask (taskId string) (Models.Task, error){
	return taskService.taskRepository.GetTaskById(taskId)
}

func (taskService *TaskService) ToggleTaskIsDone (taskId string) (Models.Task, error){
	return taskService.taskRepository.ToggleTaskIsDoneById(taskId)
}

func (TaskService *TaskService) UpdateTask(task Models.Task) (Models.Task, error){
	return TaskService.taskRepository.UpdateTask(task)
}

func (TaskService *TaskService) Delete(taskId string) (Models.Task, error){
	return TaskService.taskRepository.DeleteTaskById(taskId)
}

func (TaskService *TaskService) DeleteTasksListId (listId string) ([]Models.Task, error){
	return TaskService.taskRepository.DeleteTasksByListId(listId)
}

