package endpoints

import (
	"context"
	"fmt"
	"github.com/offluck/abgorrence/server"
)

type TestRequest struct {
	Number int `json:"number"`
}

func (TestRequest) Validate() error {
	return nil
}

func (TestRequest) Type() server.RequestType {
	return server.Filled
}

type TestResponse struct {
	Message string `json:"message"`
}

func TestHandle(ctx context.Context, request TestRequest) (TestResponse, error) {
	return TestResponse{
		Message: fmt.Sprintf("Number: %d", request.Number),
	}, nil
}
