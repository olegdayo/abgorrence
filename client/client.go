package client

import (
	"fmt"
	"github.com/offluck/abgorrence/common/models/endpoint"
	"github.com/offluck/abgorrence/common/models/relations"
	"github.com/offluck/abgorrence/server"
	"log"
)

type Client struct {
	relations relations.Relations
	current   endpoint.Endpoint
	addr      string
}

func New(gateway endpoint.Endpoint, addr string) *Client {
	return &Client{
		relations: relations.New(),
		current:   gateway,
		addr:      addr,
	}
}

func FromServer(server *server.Server, gateway endpoint.Endpoint) *Client {
	return &Client{
		relations: server.GetRelations(),
		current:   gateway,
		addr:      server.Addr,
	}
}

func (c *Client) AddRelation(from endpoint.Endpoint, to endpoint.Endpoint) {
	c.relations.Add(from, to)
}

func (c *Client) DeleteRelation(from endpoint.Endpoint, to endpoint.Endpoint) {
	c.relations.Delete(from, to)
}

func (c *Client) IsRelationPresent(from endpoint.Endpoint, to endpoint.Endpoint) bool {
	return c.relations.IsRelationPresent(from, to)
}

func (c *Client) CanProceed(to endpoint.Endpoint) bool {
	return c.IsRelationPresent(c.current, to)
}

func (c *Client) SendRequest(to endpoint.Endpoint, request server.Request) error {
	if !c.CanProceed(to) {
		return fmt.Errorf("cannot get from %s to %s", c.current, to)
	}

	log.Printf("Proceed from %s to %s", c.current, to)
	c.current = to

	switch to.Method {
	case endpoint.GET:
	case endpoint.POST:
	case endpoint.PUT:
	case endpoint.PATCH:
	case endpoint.DELETE:
	}

	return nil
}
