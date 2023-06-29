package main

import (
	"github.com/offluck/abgorrence/common/models/endpoint"
	"github.com/offluck/abgorrence/common/models/endpoint/datatype"
	"github.com/offluck/abgorrence/examples/simple-server/endpoints"
	"github.com/offluck/abgorrence/server"
)

func main() {
	s := server.New(":8080")
	baseURL := "http://localhost:8080"

	gateway := endpoint.New(
		"gateway",
		baseURL,
		"/",
		endpoint.GET,
		datatype.ApplicationJSON,
	)

	test := endpoint.New(
		"test",
		baseURL,
		"/test",
		endpoint.GET,
		datatype.ApplicationJSON,
	)

	s.AddRelation(
		gateway,
		test,
	)

	s.AddHandler(
		gateway,
		server.Init(endpoints.GateWayHandle),
	)
	s.AddHandler(
		test,
		server.Init(endpoints.TestHandle),
	)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
