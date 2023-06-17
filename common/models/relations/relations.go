package relations

import (
	"github.com/offluck/abgorrence/common/models/endpoint"
	"github.com/offluck/abgorrence/common/models/set"
)

type Relations struct {
	relations map[endpoint.Endpoint]*set.Set[endpoint.Endpoint]
}

func New() Relations {
	return Relations{
		relations: make(map[endpoint.Endpoint]*set.Set[endpoint.Endpoint]),
	}
}

func (r *Relations) Add(from endpoint.Endpoint, to endpoint.Endpoint) {
	if r.relations[from] == nil {
		r.relations[from] = set.New[endpoint.Endpoint]()
	}
	r.relations[from].Add(to)
}

func (r *Relations) Delete(from endpoint.Endpoint, to endpoint.Endpoint) {
	if r.relations[from] == nil {
		return
	}
	r.relations[from].Delete(to)
}

func (r *Relations) IsRelationPresent(from endpoint.Endpoint, to endpoint.Endpoint) bool {
	if r.relations[from] == nil {
		return false
	}
	return r.relations[from].Contains(to)
}

func (r *Relations) GetRelationsFor(from endpoint.Endpoint) []endpoint.Endpoint {
	if r.relations[from] == nil {
		return []endpoint.Endpoint{}
	}
	return r.relations[from].GetAll()
}
