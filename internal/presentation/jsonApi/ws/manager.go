package ws

import (
	"errors"
	"net/http"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/auth"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/ws/client"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/ws/handler"
	"github.com/gorilla/websocket"
)

// Manager is used to hold references to all Clients
type Manager struct {
	authManager     auth.Manager
	roomMemberQuery query.IRoomMember

	clients *client.Clients
}

func NewManager(
	authManager auth.Manager,
) *Manager {
	return &Manager{
		authManager: authManager,

		clients: client.NewClients(),
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	// Create a new websocket connection
	client, err := client.NewClient(w, r, m.authManager)
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
	defer m.clients.Remove(client)

	global.Log().Info("WS_CLIENT_CONNECTED")

	errorHandler := handler.NewError(client.GetID(), m.clients)
	for {
		var e error

		select {
		case event := <-client.Read():
			if err := routeEvent(event,
				handler.NewRoomMessage(client.GetID(), m.clients, m.roomMemberQuery),
			); err != nil {
				e = err
			}
		case err := <-client.Err():
			e = err
		}

		if e != nil {
			if e, ok := err.(*websocket.CloseError); ok {
				switch e.Code {
				case websocket.CloseNormalClosure:
					global.Log().Info("WS_CLIENT_CLOSE_NORMAL")
				default:
					global.Log().Error("WS_CLIENT_CLOSE_ERROR", "code", e.Code, "text", e.Text)
				}
				continue
			}

			switch coderr.ErrorCode(e) {
			case coderr.CodeInternal:
				errorHandler.HandleError(errors.New("internal error"), client)
				global.Log().Error("WS_INTERNAL_ERROR", "err", e)
			default:
				errorHandler.HandleError(e, client)
			}
		}
	}
}
