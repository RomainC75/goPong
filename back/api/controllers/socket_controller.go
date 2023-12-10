package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	Services "github.com/saegus/test-technique-romain-chenard/api/services"

	Manager "github.com/saegus/test-technique-romain-chenard/api/sockets/managers"
)

type Controller struct {
	taskService Services.TaskServiceInterface
	listService Services.ListServiceInterface
	manager Manager.ManagerInterface
}

func New() *Controller {
	return &Controller{
		taskService: Services.NewTaskSrv(),
		listService: Services.NewListSrv(),
		manager: Manager.New(),
	}
}

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	// CheckOrigin: func(r *http.Request) bool { return true },
// 	CheckOrigin: func(r *http.Request) bool {
//         origin := r.Header.Get("Origin")
//     	// return origin == "http://localhost:3000"
// 		cfg := configu.Get()
// 		frontUrl := cfg.Front.Host
//     	return origin == frontUrl
//     },
// }

func (controller *Controller) Socket(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userIdStr, _ := userId.(string)
	userUuid, _ := uuid.Parse(userIdStr)
	userEmail, _ := c.Get("user_email")
	userEmailStr, _ := userEmail.(string)


	userData := Manager.UserData{
		UserId: userUuid, 
		UserEmail: userEmailStr,
	}

	controller.manager.ServeWS(c.Writer, c.Request, userData)
	
	
	// conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	
	// if err != nil {
	// 	log.Printf("=======> socket server configuration error: %v\n", err)
	//  	return
	// }ListService{}
	// defer conn.Close()

	// for {
	// 	var message SocketMessage.WebSocketMessage
	// 	err := conn.ReadJSON(&message)
	// 	if !errors.Is(err, nil) {
	// 		log.Printf("error occurred: %v", err)
	// 		break
	// 	}
	// 	log.Println(message)

	//  	conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!  // "+message.Content["message"]))
	//  	time.Sleep(time.Second)
	// }

	


	// Game.GameCore(conn)
}
