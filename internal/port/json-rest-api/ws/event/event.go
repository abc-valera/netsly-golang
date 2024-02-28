package event

import "encoding/json"

// Type is the type of the event
type Type string

// Event is the struct that represents the data sent over the websocket
type Event struct {
	Type    Type            `json:"type"`
	Payload json.RawMessage `json:"payload"`
}
