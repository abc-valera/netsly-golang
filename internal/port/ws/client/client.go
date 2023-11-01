package client

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn

	readChan  chan Event
	writeChan chan Event
	errChan   chan error
}

func NewClient(conn *websocket.Conn) (*Client, error) {
	client := &Client{
		connection: conn,

		readChan:  make(chan Event),
		writeChan: make(chan Event),
		errChan:   make(chan error),
	}

	go func(c *Client) {
		for {
			_, msg, err := c.connection.ReadMessage()
			if err != nil {
				c.errChan <- err
				return
			}

			var e Event
			if err := json.Unmarshal(msg, &e); err != nil {
				c.errChan <- err
				return
			}

			c.readChan <- e
		}
	}(client)

	go func(c *Client) {
		for {
			e := <-c.writeChan

			msg, err := json.Marshal(e)
			if err != nil {
				c.errChan <- err
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, msg); err != nil {
				c.errChan <- err
				return
			}
		}
	}(client)

	return client, nil
}

func (c *Client) Read() <-chan Event {
	return c.readChan
}

func (c *Client) Write() chan<- Event {
	return c.writeChan
}

func (c *Client) Err() <-chan error {
	return c.errChan
}
