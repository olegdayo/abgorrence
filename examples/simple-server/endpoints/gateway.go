package endpoints

import (
	"context"
	"github.com/offluck/abgorrence/server"
)

type GatewayRequest struct {
}

func (GatewayRequest) Validate() error {
	return nil
}

func (GatewayRequest) Type() server.RequestType {
	return server.Empty
}

type GatewayResponse struct {
	Message string `json:"message"`
}

type Handler struct {
}

func GateWayHandle(ctx context.Context, request GatewayRequest) (GatewayResponse, error) {
	return GatewayResponse{
		Message: "This is gateway",
	}, nil
}
