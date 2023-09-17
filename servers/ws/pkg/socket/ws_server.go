package socket

import (
	log "github.com/sirupsen/logrus"
)

type WsServer struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	Rooms      map[*Room]bool
}

func NewWsServer() *WsServer {
	return &WsServer{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
		Rooms:      make(map[*Room]bool),
	}
}

func (wsServer *WsServer) Start() {
	for {
		select {
		case client := <-wsServer.Register:
			wsServer.registerClient(client)
		case client := <-wsServer.Unregister:
			wsServer.unregisterClient(client)
		case message := <-wsServer.Broadcast:
			wsServer.broadcastToClients(message)
		}
	}
}

func (wsServer *WsServer) registerClient(client *Client) {
	wsServer.Clients[client] = true
}

func (wsServer *WsServer) unregisterClient(client *Client) {
	if _, ok := wsServer.Clients[client]; ok {
		delete(wsServer.Clients, client)
	}
}

func (wsServer *WsServer) broadcastToClients(message []byte) {
	for client := range wsServer.Clients {
		client.send <- message
	}
}

func (wsServer *WsServer) FindRoomByID(roomID string) *Room {
	var foundRoom *Room
	for room := range wsServer.Rooms {
		if room.RoomID == roomID {
			foundRoom = room
			break
		}
	}
	return foundRoom
}

func (WsServer *WsServer) FindClientByID(clientID string) *Client {
	var foundClient *Client
	for client := range WsServer.Clients {
		if client.ClientID == clientID {
			foundClient = client
			break
		}
	}
	return foundClient
}

func (wsServer *WsServer) CreateRoom(private bool, maxPlayers int) *Room {
	room := NewRoom(private, maxPlayers)
	go room.Start()
	wsServer.Rooms[room] = true
	log.Info("Room created with ID: ", room.RoomID)
	log.Info("Number of rooms created: ", len(wsServer.Rooms))
	return room
}
