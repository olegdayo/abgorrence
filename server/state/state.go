package state

import (
	"fmt"
	"github.com/offluck/abgorrence/server/endpoint"
	"log"
)

type State struct {
	rels    relations
	current endpoint.Endpoint
}

func New(gateway endpoint.Endpoint) *State {
	return &State{
		rels:    create(),
		current: gateway,
	}
}

func (s *State) AddRelation(from endpoint.Endpoint, to endpoint.Endpoint) {
	s.rels.add(from, to)
}

func (s *State) DeleteRelation(from endpoint.Endpoint, to endpoint.Endpoint) {
	s.rels.delete(from, to)
}

func (s *State) IsRelationPresent(from endpoint.Endpoint, to endpoint.Endpoint) bool {
	return s.rels.isRelationPresent(from, to)
}

func (s *State) CanProceed(to endpoint.Endpoint) bool {
	return s.IsRelationPresent(s.current, to)
}

func (s *State) Proceed(to endpoint.Endpoint) error {
	if !s.CanProceed(to) {
		return fmt.Errorf("cannot get from %s to %s", s.current, to)
	}

	log.Printf("Proceed from %s to %s", s.current, to)
	s.current = to
	return nil
}
