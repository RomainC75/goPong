package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	ListService "github.com/saegus/test-technique-romain-chenard/internal/modules/list/services"
	TaskService "github.com/saegus/test-technique-romain-chenard/internal/modules/task/services"

	Game "github.com/saegus/test-technique-romain-chenard/internal/modules/game/core"
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

// cfg := configu.Get()
// frontUrl := cfg.Front.Host

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin: func(r *http.Request) bool {
    //     origin := r.Header.Get("Origin")
    // 	return origin == "http://localhost:3000"
    // },
}

func (controller *Controller) Socket(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
	 return
	}
	// defer conn.Close()

	// for {
	// 	var message SocketMessage.Message
	// 	err := conn.ReadJSON(&message)
	// 	if !errors.Is(err, nil) {
	// 		log.Printf("error occurred: %v", err)
	// 		break
	// 	}
	// 	log.Println(message)

	//  	conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
	//  	time.Sleep(time.Second)
	// }

	Game.GameCore(conn)
}
