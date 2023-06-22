package main

import (
	"github.com/offluck/abgorrence/common/models/endpoint"
	"github.com/offluck/abgorrence/server"
)

func main() {
	serv := server.New()
	serv.AddRelation(
		endpoint.New("/", endpoint.GET),
		endpoint.New("/test", endpoint.GET),
	)
	serv.Addr = ":8080"
	serv.ListenAndServe()
}
