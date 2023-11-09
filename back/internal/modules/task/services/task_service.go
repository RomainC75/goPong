package services

import (
	"github.com/google/uuid"
	"github.com/saegus/test-technique-romain-chenard/internal/modules/task/models"
	TaskModel "github.com/saegus/test-technique-romain-chenard/internal/modules/task/models"
	TaskRepository "github.com/saegus/test-technique-romain-chenard/internal/modules/task/repositories"
	TaskRequest "github.com/saegus/test-technique-romain-chenard/internal/modules/task/requests"
)

type TaskService struct {
	taskRepository TaskRepository.TaskRepositoryInterface
}

func New() *TaskService{
	return &TaskService{
		taskRepository: TaskRepository.New(),
	}
}

func (taskService *TaskService) CreateTask (task TaskRequest.CreateTaskRequest, listId string) (TaskModel.Task, error){
	var newTask TaskModel.Task
	listUuid := uuid.MustParse(listId)

	newTask.Name = task.Name
	newTask.ListId = listUuid
	newTask.DeadLine = task.DeadLine
	newTask.Description = task.Description
	newTask.IsDone = false
	
	newTask, err := taskService.taskRepository.CreateTask(newTask)

	if err != nil{
		return TaskModel.Task{}, err
	}
	return newTask, nil
}


func (taskService *TaskService) GetTasks (listId string) []models.Task{
	return taskService.taskRepository.GetTasks(listId)
}

func (taskService *TaskService) GetTask (taskId string) (models.Task, error){
	return taskService.taskRepository.GetTaskById(taskId)
}

func (taskService *TaskService) ToggleTaskIsDone (taskId string) (models.Task, error){
	return taskService.taskRepository.ToggleTaskIsDoneById(taskId)
}

func (TaskService *TaskService) UpdateTask(task models.Task) (models.Task, error){
	return TaskService.taskRepository.UpdateTask(task)
}

func (TaskService *TaskService) Delete(taskId string) (models.Task, error){
	return TaskService.taskRepository.DeleteTaskById(taskId)
}

func (TaskService *TaskService) DeleteTasksListId (listId string) ([]models.Task, error){
	return TaskService.taskRepository.DeleteTasksByListId(listId)
}

