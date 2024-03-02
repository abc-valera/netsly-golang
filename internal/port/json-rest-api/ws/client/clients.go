package client

import "sync"

type Clients struct {
	clients map[string](map[Client]bool)
	sync.RWMutex
}

func NewClients() *Clients {
	return &Clients{
		clients: make(map[string]map[Client]bool),
	}
}

func (c *Clients) Add(client Client) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.clients[client.GetID()]; !ok {
		c.clients[client.GetID()] = make(map[Client]bool)
	}
	c.clients[client.GetID()][client] = true
}

func (c *Clients) Remove(client Client) {
	c.Lock()
	defer c.Unlock()

	// Check if connection exists
	if _, ok := c.clients[client.GetID()][client]; ok {
		client.close()
		delete(c.clients[client.GetID()], client)
	}
}

func (c *Clients) GetAll() []Client {
	c.RLock()
	defer c.RUnlock()

	clients := make([]Client, 0, len(c.clients))
	for _, clientsMap := range c.clients {
		for client := range clientsMap {
			clients = append(clients, client)
		}
	}

	return clients
}

func (c *Clients) GetAllByUserID(userID string) []Client {
	c.RLock()
	defer c.RUnlock()

	clients := make([]Client, 0, len(c.clients[userID]))
	for client := range c.clients[userID] {
		clients = append(clients, client)
	}

	return clients
}
