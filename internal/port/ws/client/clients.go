package client

import "sync"

type Clients struct {
	clients map[*Client]bool
	sync.RWMutex
}

func NewClients() *Clients {
	return &Clients{
		clients: make(map[*Client]bool),
	}
}

func (c *Clients) Add(client *Client) {
	c.Lock()
	defer c.Unlock()

	c.clients[client] = true
}

func (c *Clients) Remove(client *Client) {
	c.Lock()
	defer c.Unlock()

	// Check if Client exists, then delete it
	if _, ok := c.clients[client]; ok {
		// close connection
		client.connection.Close()
		delete(c.clients, client)
	}
}
