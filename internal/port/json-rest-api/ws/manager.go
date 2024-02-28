package ws

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/abc-valera/netsly-api-golang/internal/domain"
	"github.com/abc-valera/netsly-api-golang/internal/domain/coderr"
	"github.com/abc-valera/netsly-api-golang/internal/domain/global"
	"github.com/abc-valera/netsly-api-golang/internal/domain/service"
	"github.com/abc-valera/netsly-api-golang/internal/port/json-rest-api/ws/event"
	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	readLimit int64 = 1024

	pongWait     = 5 * time.Second
	pingInterval = 3 * time.Second
)

// Manager is used to hold references to all Clients
type Manager struct {
	tokenMaker service.ITokenMaker

	conns *connections
}

func NewManager(
	services domain.Services,
) *Manager {
	return &Manager{
		tokenMaker: services.TokenMaker,

		conns: newConnections(),
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	// Authenticate the client, if the token is invalid, return 401
	authPayload, err := m.tokenMaker.VerifyToken(r.URL.Query().Get("token"))
	if err != nil {
		switch coderr.ErrorCode(err) {
		case coderr.CodeUnauthenticated, coderr.CodeInvalidArgument:
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		default:
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	fmt.Println(authPayload)

	// Upgrade the HTTP connection to a websocket connection
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		global.Log().Error("WS_UPGRADE_ERROR", "err", coderr.NewInternal(err))
		return
	}
	m.conns.Add(conn)
	global.Log().Info("WS_CLIENT_CONNECTED")

	// Deferring the removal of the connection
	defer func() {
		m.conns.Remove(conn)
		global.Log().Info("WS_CLIENT_DISCONNECTED")
	}()

	writeChan := make(chan event.Event)
	readChan := make(chan event.Event)
	errChan := make(chan error) // errChan is used for other client errors

	pingChan := make(chan struct{})

	conn.SetReadLimit(readLimit)

	// Send pings asynchronously to the pingChan
	go func() {
		ticker := time.NewTicker(pingInterval)
		defer ticker.Stop()

		for {
			<-ticker.C
			pingChan <- struct{}{}
		}
	}()

	// Set initial connection read deadline
	if err := conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		global.Log().Error("WS_CLIENT_ERROR", "err", coderr.NewInternal(err))
		return
	}
	// Set Pong handler to update the read deadline
	conn.SetPongHandler(func(appData string) error {
		fmt.Println("Received pong")
		return conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	// Write to the websocket asynchronously
	go func() {
		for {
			select {
			case e := <-writeChan:
				msg, err := json.Marshal(e)
				if err != nil {
					errChan <- coderr.NewInternal(err)
					return
				}

				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					errChan <- coderr.NewInternal(err)
					return
				}
			case <-pingChan:
				fmt.Println("Sending ping")
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					errChan <- coderr.NewInternal(err)
					return
				}
			}
		}
	}()

	// Read from the websocket asynchronously
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				errChan <- coderr.NewInternal(err)
				return
			}

			var e event.Event
			if err := json.Unmarshal(msg, &e); err != nil {
				errChan <- coderr.NewMessage(coderr.CodeInvalidArgument, "json input syntax error")
				return
			}

			readChan <- e
		}
	}()

	// Handle the different channels output
	for {
		select {
		case msg := <-readChan:
			if err := routeEvent(msg, conn); err != nil {
				// Handle handler errors here
				switch coderr.ErrorCode(err) {
				}
				global.Log().Error("WS_HANDLER_ERROR", err)
			}
		case err := <-errChan:
			if e, ok := err.(*websocket.CloseError); ok {
				// Handle WS close errors here
				switch e.Code {
				}
			} else {
				global.Log().Error("WS_CLIENT_ERROR", err)
			}
		}
	}
}

// connections is a thread-safe map of websocket connections
type connections struct {
	conns map[*websocket.Conn]bool
	sync.RWMutex
}

func newConnections() *connections {
	return &connections{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (c *connections) Add(conn *websocket.Conn) {
	c.Lock()
	defer c.Unlock()

	c.conns[conn] = true
}

func (c *connections) Remove(conn *websocket.Conn) {
	c.Lock()
	defer c.Unlock()

	// Check if connection exists, then delete it
	if _, ok := c.conns[conn]; ok {
		// close connection
		conn.Close()
		delete(c.conns, conn)
	}
}
