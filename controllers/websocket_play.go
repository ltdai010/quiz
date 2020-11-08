package controllers

import (
	"net/http"
	"taskmanagerproject/websocket/arrangetask/broadcastbody"
)

var (
	clients  = make(map[*websocket.Conn]string)   // connected clients
	Bc       = make(chan broadcastbody.Broadcast) // broadcastbody channel
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
