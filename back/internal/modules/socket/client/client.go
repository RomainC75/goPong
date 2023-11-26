package client

import (
	"github.com/gorilla/websocket"
	Manager "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/manager"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager *Manager.Manager
}

func NewClient(conn *websocket.Conn, manager *Manager.Manager) *Client{
	return &Client{
		connection: conn,
		manager: manager,
	}
}