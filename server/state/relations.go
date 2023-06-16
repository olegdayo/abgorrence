package state

import "github.com/offluck/abgorrence/server/endpoint"

type relations struct {
	rels map[endpoint.Endpoint]*set[endpoint.Endpoint]
}

func create() relations {
	return relations{
		rels: make(map[endpoint.Endpoint]*set[endpoint.Endpoint]),
	}
}

func (r *relations) add(from endpoint.Endpoint, to endpoint.Endpoint) {
	if r.rels[from] == nil {
		r.rels[from] = initialize[endpoint.Endpoint]()
	}
	r.rels[from].add(to)
}

func (r *relations) delete(from endpoint.Endpoint, to endpoint.Endpoint) {
	if r.rels[from] == nil {
		return
	}
	r.rels[from].delete(to)
}

func (r *relations) isRelationPresent(from endpoint.Endpoint, to endpoint.Endpoint) bool {
	if r.rels[from] == nil {
		return false
	}
	return r.rels[from].contains(to)
}
