package controllers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	ListService "github.com/saegus/test-technique-romain-chenard/internal/modules/list/services"
	TaskService "github.com/saegus/test-technique-romain-chenard/internal/modules/task/services"

	Game "github.com/saegus/test-technique-romain-chenard/internal/modules/game/core"
	configu "github.com/saegus/test-technique-romain-chenard/pkg/configu"

	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
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
	// CheckOrigin: func(r *http.Request) bool { return true },
	CheckOrigin: func(r *http.Request) bool {
        origin := r.Header.Get("Origin")
    	// return origin == "http://localhost:3000"
		cfg := configu.Get()
		frontUrl := cfg.Front.Host
    	return origin == frontUrl
    },
}

func (controller *Controller) Socket(c *gin.Context) {
	// upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	
	if err != nil {
		log.Printf("=======> socket server configuration error: %v\n", err)
	 	return
	}
	defer conn.Close()

	for {
		var message SocketMessage.WebSocketMessage
		err := conn.ReadJSON(&message)
		if !errors.Is(err, nil) {
			log.Printf("error occurred: %v", err)
			break
		}
		log.Println(message)

	 	conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
	 	time.Sleep(time.Second)
	}

	Game.GameCore(conn)
}
