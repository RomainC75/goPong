package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	TaskRequest "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	TaskResponse "github.com/saegus/test-technique-romain-chenard/api/dto/responses"
	Services "github.com/saegus/test-technique-romain-chenard/api/services"
	"github.com/saegus/test-technique-romain-chenard/data/models"
)

type TaskController struct {
	taskService Services.TaskServiceInterface
	listService Services.ListServiceInterface
}

func NewTaskCtrl() *TaskController {
	return &TaskController{
		taskService: Services.NewTaskSrv(),
		listService: Services.NewListSrv(),
	}
}

func (controller *TaskController) CreateTask(c *gin.Context) {
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

func (controller *TaskController) GetTasks(c *gin.Context) {
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

func (controller *TaskController) ToogleTask(c *gin.Context) {
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

func (controller *TaskController) UpdateTask(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)

	taskId := c.Param("taskId")
	var newTask models.Task
	if err := c.ShouldBind(&newTask); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

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

func (controller *TaskController) DeleteTask(c *gin.Context) {
	taskId := c.Param("taskId")

	deletedTask, err := controller.taskService.Delete(taskId)
	if  err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, TaskResponse.ToTaskResponse(deletedTask))
	
}