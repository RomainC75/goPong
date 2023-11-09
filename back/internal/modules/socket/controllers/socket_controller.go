package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	ListService "github.com/saegus/test-technique-romain-chenard/internal/modules/list/services"
	TaskService "github.com/saegus/test-technique-romain-chenard/internal/modules/task/services"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
   }

func (controller *Controller) Socket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
	 return
	}
	defer conn.Close()
	for {
	 conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
	 time.Sleep(time.Second)
	}
}
