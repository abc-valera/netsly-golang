package manager

import (
	"log"
	"net/http"

	"github.com/abc-valera/flugo-api-golang/internal/port/ws/client"
	"github.com/abc-valera/flugo-api-golang/internal/port/ws/handler"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
)

// Manager is used to hold references to all Clients
type Manager struct {
	clients *client.Clients
}

func NewManager() *Manager {
	return &Manager{
		clients: client.NewClients(),
	}
}

// ServeWS is a HTTP Handler that the has the Manager that allows connections
func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client, err := client.NewClient(conn)
	if err != nil {
		return
	}
	m.clients.Add(client)

	defer m.clients.Remove(client)

	for {
		select {
		case msg := <-client.Read():
			if err := m.routeEvent(msg, client); err != nil {
				log.Println("Error: ", err)
				return
			}
		case err := <-client.Err():
			log.Println("Error: ", err)
			return
		}
	}
}

func (m *Manager) routeEvent(e client.Event, c *client.Client) error {
	for {
		switch e.Type {
		case handler.EventTypeSendMessage:
			if err := handler.SendChatMsgHandler(e, c); err != nil {
				return err
			}
		default:
			return client.ErrEventNotSupported
		}
	}
}
