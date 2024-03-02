package ws

import (
	"net/http"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/client"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/handler"
	"github.com/gorilla/websocket"
)

// Manager is used to hold references to all Clients
type Manager struct {
	tokenMaker      service.ITokenMaker
	roomMemberQuery query.IRoomMember

	clients *client.Clients
}

func NewManager(
	services domain.Services,
) *Manager {
	return &Manager{
		tokenMaker: services.TokenMaker,

		clients: client.NewClients(),
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	// Create a new websocket connection
	client, err := client.NewClient(w, r, m.tokenMaker)
	if err != nil {
		switch coderr.ErrorCode(err) {
		case coderr.CodeUnauthenticated, coderr.CodeInvalidArgument:
			global.Log().Error("WS_CLIENT_AUTH_ERROR", "err", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
		default:
			global.Log().Error("WS_CLIENT_ERROR", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	m.clients.Add(client)
	defer m.clients.Remove(client) // Defer removal of the client

	global.Log().Info("WS_CLIENT_CONNECTED")

	for {
		select {
		case e := <-client.Read():
			if err := routeEvent(e,
				handler.NewError(client.GetID(), m.clients),
				handler.NewRoom(client.GetID(), m.clients, m.roomMemberQuery),
			); err != nil {
				switch coderr.ErrorCode(err) {
				case coderr.CodeInvalidArgument:
					global.Log().Error("WS_CLIENT_401_ERROR", "err", err)
					http.Error(w, err.Error(), http.StatusBadRequest)
				default:
					global.Log().Error("WS_CLIENT_ERROR", "err", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				global.Log().Error("WS_HANDLER_ERROR", "err", err)
			}
		case err := <-client.Err():
			if e, ok := err.(*websocket.CloseError); ok {
				// Handle WS close errors here
				switch e.Code {
				}
				global.Log().Error("WS_CLIENT_CLOSE_ERROR", "code", e.Code, "text", e.Text)
				continue
			}

			switch coderr.ErrorCode(err) {
			case coderr.CodeInvalidArgument:
				global.Log().Error("WS_CLIENT_401_ERROR", "err", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				global.Log().Error("WS_CLIENT_ERROR", "err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
