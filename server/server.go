package server

import (
	"AgarGo/server/managers"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Server struct {
	clients   map[uint32]*Client
	upgrader  *websocket.Upgrader
	idManager *managers.IdManager
}

func NewServer() *Server {
	return &Server{
		clients:   make(map[uint32]*Client),
		upgrader:  &websocket.Upgrader{},
		idManager: managers.NewIdManager(),
	}
}

func (s *Server) Listen() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		client := NewClient(s.idManager.GenerateClientId(), conn)

		s.clients[client.id] = client
	}

	http.HandleFunc("/", handler)
}
