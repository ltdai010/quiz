package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	clients = make(map[*websocket.Conn]bool) // connected clients
	broadcast = make(chan ScoreBoard)           // broadcast channel
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type ScoreBoard struct {
	scoreList map[string]int
	playerList map[string]string
}

type Answer struct {
	choose int
	Time int64
}

func main()  {
	
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	// Register our new client
	clients[ws] = true
	for {
		var ans Answer
		var sb ScoreBoard
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&ans)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		calculateScore(sb)
		// Send the score board to client
		broadcast <- sb
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		sb := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(sb)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func calculateScore(sb ScoreBoard) ScoreBoard {
	return sb
}