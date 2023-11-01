package client

import (
	"encoding/json"
	"errors"
)

// Event is the struct that represents the data sent over the websocket
type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

var (
	ErrEventNotSupported = errors.New("Event type is not supported")
)

type EventHandler func(event Event, c *Client) error
