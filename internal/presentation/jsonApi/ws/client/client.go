package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abc-valera/netsly-golang/internal/core/coderr"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/auth"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/ws/event"
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

// Client is a decorator for a websocket connection.
// It provides a way to read and write events to the connection.
type Client interface {
	GetID() string

	Read() <-chan event.Event
	Write() chan<- event.Event
	Err() <-chan error

	close() error
}

// client is a concrete implementation of the Client interface.
type client struct {
	connection *websocket.Conn

	readChan  chan event.Event
	writeChan chan event.Event
	errChan   chan error
	pingChan  chan string

	userID string
}

func NewClient(w http.ResponseWriter, r *http.Request, authManager auth.Manager) (Client, error) {
	// Authenticate the client, if the token is invalid, return 401
	authPayload, err := authManager.VerifyToken(r.URL.Query().Get("token"))
	if err != nil {
		switch coderr.ErrorCode(err) {
		case coderr.CodeUnauthenticated, coderr.CodeInvalidArgument:
			return nil, err
		default:
			return nil, err
		}
	}

	// Upgrade the HTTP connection to a websocket connection
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, coderr.NewCodeError(coderr.CodeInvalidArgument, err)
	}

	client := &client{
		connection: conn,

		readChan:  make(chan event.Event),
		writeChan: make(chan event.Event),
		errChan:   make(chan error),
		pingChan:  make(chan string),

		userID: authPayload.UserID,
	}

	// Set the maximum read limit
	conn.SetReadLimit(readLimit)

	// Set initial connection read deadline
	if err := conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		client.errChan <- err
	}

	// Send pings asynchronously to the pingChan
	go func() {
		ticker := time.NewTicker(pingInterval)
		defer ticker.Stop()

		for {
			<-ticker.C
			client.pingChan <- "ping"
		}
	}()

	// Set Pong handler to update the read deadline
	conn.SetPongHandler(func(appData string) error {
		return conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	// Read from the websocket asynchronously
	go func() {
		for {
			_, msg, err := client.connection.ReadMessage()
			if err != nil {
				client.errChan <- err
			}

			var e event.Event
			if err := json.Unmarshal(msg, &e); err != nil {
				client.errChan <- coderr.NewCodeMessage(coderr.CodeInvalidArgument, "json input syntax error")
			}

			client.readChan <- e
		}
	}()

	// Write to the websocket asynchronously
	go func() {
		for {
			select {
			case e := <-client.writeChan:
				msg, err := json.Marshal(e)
				if err != nil {
					client.errChan <- coderr.NewInternalErr(err)
				}

				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					client.errChan <- err
				}
			case <-client.pingChan:
				if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					client.errChan <- err
				}
			}
		}
	}()

	return client, nil
}

func (c client) GetID() string {
	return c.userID
}

func (c client) Read() <-chan event.Event {
	return c.readChan
}

func (c client) Write() chan<- event.Event {
	return c.writeChan
}

func (c client) Err() <-chan error {
	return c.errChan
}

func (c client) close() error {
	close(c.readChan)
	close(c.writeChan)
	close(c.errChan)
	close(c.pingChan)

	return c.connection.Close()
}
