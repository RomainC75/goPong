package manager

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
	configu "github.com/saegus/test-technique-romain-chenard/pkg/configu"

	// SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
	Client "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/client"
)

type ManagerInterface interface{
	serveWS(w gin.ResponseWriter , r *http.Request)
}

var(
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			// return origin == "http://localhost:3000"
			cfg := configu.Get()
			frontUrl := cfg.Front.Host
			return origin == frontUrl
		},
	}
)

type Manager struct {
	client Client.ClientList
	sync.RWMutex
}

func NewManager() *Manager{
	return &Manager{
		client: make(Client.ClientList),
	}
}

func (m *Manager) serveWS(w gin.ResponseWriter , r *http.Request){
	log.Println("new Connection")
	conn, err := websocketUpgrader.Upgrade(w,r, nil)
	if err != nil{
		log.Println(err)
		return
	}
	conn.Close()
}