package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ListService "github.com/saegus/test-technique-romain-chenard/internal/modules/list/services"
	"github.com/saegus/test-technique-romain-chenard/internal/modules/task/models"
	TaskRequest "github.com/saegus/test-technique-romain-chenard/internal/modules/task/requests"
	TaskResponse "github.com/saegus/test-technique-romain-chenard/internal/modules/task/responses"
	TaskService "github.com/saegus/test-technique-romain-chenard/internal/modules/task/services"
	"github.com/saegus/test-technique-romain-chenard/pkg/utils"
)

type Controller struct {
	taskService TaskService.TaskServiceInterface
	listService ListService.ListServiceInterface
}

func New() *Controller {
	return &Controller{
		taskService: TaskService.New(),
		listService: ListService.New(),
	}
}

func (controller *Controller) CreateTask(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)

	var newTask TaskRequest.CreateTaskRequest
	if err := c.ShouldBind(&newTask); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// s1, _:= time.Parse(time.RFC3339, “2018-12-12”)
	// parsedDate, _ := time.Parse("2006-01-01T00:00:00Z", newTask.DeadLine.String() )
	// newTask.DeadLine = parsedDate

	listId := c.Param("listId")
	foundList, err := controller.listService.GetList(listId)
	if  err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}


	if userIdStr != foundList.UserId.String(){
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "not authorized to modify this list"})
		return
	}

	recordedList, err := controller.taskService.CreateTask(newTask, listId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recordedList)
}

func (controller *Controller) GetTasks(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)

	var newTask models.Task
	if err := c.ShouldBind(&newTask); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	listId := c.Param("listId")
	foundList, err := controller.listService.GetList(listId)
	
	if  err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	
	if userIdStr != foundList.UserId.String(){
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "not authorized to modify this list"})
		return
	}

	c.JSON(http.StatusOK, TaskResponse.ToTaskArrayResponse(controller.taskService.GetTasks(listId)))
}

func (controller *Controller) ToogleTask(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)

	taskId := c.Param("taskId")

	// test task
	foundTask, err := controller.taskService.GetTask(taskId)
	
	if  err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// test list

	foundList, err := controller.listService.GetList(foundTask.ListId.String())
	
	if  err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	
	if userIdStr != foundList.UserId.String(){
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "not authorized to modify this task"})
		return
	}

	//
	
	newTask, err := controller.taskService.ToggleTaskIsDone(taskId)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, newTask)
}

func (controller *Controller) UpdateTask(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)

	taskId := c.Param("taskId")
	var newTask models.Task
	if err := c.ShouldBind(&newTask); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	utils.PrettyDisplay(newTask)

	// test task 
	foundTask, err := controller.taskService.GetTask(taskId)
	if  err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// test list

	isUserTheOwner, err := controller.listService.IsUserTheOwnerOfTHeList(userIdStr, foundTask.ListId.String())
	if  err != nil || !isUserTheOwner{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unauthorized to change this list"})
		return
	}

	updatedTask, err := controller.taskService.UpdateTask(newTask)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, TaskResponse.ToTaskResponse(updatedTask))
}

func (controller *Controller) DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")

	deletedTask, err := controller.taskService.Delete(taskId)
	if  err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, TaskResponse.ToTaskResponse(deletedTask))
	
}