package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ListRequest "github.com/saegus/test-technique-romain-chenard/api/dto/requests"
	ListResponse "github.com/saegus/test-technique-romain-chenard/api/dto/responses"
	Services "github.com/saegus/test-technique-romain-chenard/api/services"
	"github.com/saegus/test-technique-romain-chenard/data/models"
)

type ListController struct {
	listService Services.ListServiceInterface
	taskService Services.TaskServiceInterface
}

func NewList() *ListController {
	return &ListController{
		listService: Services.NewListSrv(),
		taskService: Services.NewTaskSrv(),
	}
}

func (controller *ListController) CreateList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)

	var newList ListRequest.CreateListRequest
	if err := c.ShouldBind(&newList); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	recordedList, err := controller.listService.CreateList(newList, userIdStr)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recordedList)
}

func (controller *ListController) GetLists(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)
	
	lists, err := controller.listService.GetLists(userIdStr)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ListResponse.ToListArrayResponse(lists))
}

func (controller *ListController) DeleteList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)
	listId := c.Param("listId")

	isOwner, err := controller.listService.IsUserTheOwnerOfTHeList(userIdStr, listId)
	if err != nil || !isOwner{
		c.JSON(http.StatusBadRequest, gin.H{"error": "user cannot delete this list and this content"})
		return
	}

	_, err = controller.taskService.DeleteTasksListId(listId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedList, err := controller.listService.DeleteList(userIdStr, listId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ListResponse.ToListResponse(deletedList))
}

func (controller *ListController) UpdateList(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)

	var updateList models.List
	if err := c.ShouldBind(&updateList); err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	updatedList, err := controller.listService.UpdateList(userIdStr, updateList)

	if err != nil{
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedList)
}
