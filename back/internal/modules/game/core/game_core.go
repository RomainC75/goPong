package game

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"

	SocketMessage "github.com/saegus/test-technique-romain-chenard/internal/modules/socket/requests"
)

type GameList map[*Game]bool

type Command struct{
	p1 string
	p2 string
}


type Game struct {
	Ball Ball
}

type Ball struct {
	Position Position
	Direction float64
}

type Position struct{
	x float64
	y float64
}



func GameCore(conn *websocket.Conn){

	go func (){
		for{
			// commands := <- commandIn
			// fmt.Printf("command", commands)

			var message SocketMessage.Message
			err := conn.ReadJSON(&message)
			if !errors.Is(err, nil) {
				log.Printf("error occurred: %v", err)
				break
			}

			fmt.Println("MESSAGEZ : ", message.Message)

			conn.WriteMessage(websocket.TextMessage, []byte(message.Message))
			time.Sleep(time.Millisecond * 500)
		}
		defer conn.Close()
	}()
}