package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/offluck/abgorrence/common/models/endpoint"
	"log"
	"net/http"
)

type handlerWrapper interface {
	http.Handler
	SetSource(endpoint.Endpoint)
	SetTargets([]endpoint.Endpoint)
}

type Wrapper[Req Request, Resp Response] struct {
	handler func(ctx context.Context, req Req) (Resp, error)
	source  endpoint.Endpoint
	targets []endpoint.Endpoint
}

func Init[Req Request, Resp Response](handler func(ctx context.Context, req Req) (Resp, error)) *Wrapper[Req, Resp] {
	return &Wrapper[Req, Resp]{
		handler: handler,
	}
}

func (w *Wrapper[Req, Resp]) SetSource(source endpoint.Endpoint) {
	w.source = source
}

func (w *Wrapper[Req, Resp]) SetTargets(targets []endpoint.Endpoint) {
	w.targets = targets
}

func (w *Wrapper[Req, Resp]) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != w.source.Method.ToString() {
		rw.WriteHeader(http.StatusBadRequest)
		writeError(rw, "bad request", fmt.Errorf("expected method: %s, got method: %s", w.source.Method.ToString(), r.Method))
		return
	}

	var req Req
	if req.Type() == Filled {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			writeError(rw, "failed to parse request", err)
			return
		}

		errValidation := req.Validate()
		if errValidation != nil {
			rw.WriteHeader(http.StatusBadRequest)
			writeError(rw, "bad request", errValidation)
			return
		}
	}

	resp, err := w.handler(r.Context(), req)
	if err != nil {
		log.Printf("execution fail: %s", err)
		rw.WriteHeader(http.StatusInternalServerError)
		writeError(rw, "execute handler", err)
		return
	}

	rawData, err := json.Marshal(&resp)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		writeError(rw, "decode response", err)
		return
	}

	for _, target := range w.targets {
		rw.Header().Add("link", target.GetRelation())
	}

	_, _ = rw.Write(rawData)
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

func newErrorResponse(message string, err error) ErrorResponse {
	return ErrorResponse{
		ErrorMessage: fmt.Sprintf("%s: %+v", message, err),
	}
}

func writeError(rw http.ResponseWriter, message string, err error) {
	log.Printf("Writing error. %s: %+v", message, err)
	bytes, err := json.Marshal(newErrorResponse(message, err))
	if err != nil {
		log.Printf("Error during writing: %+v", err)
	}

	n, err := rw.Write(bytes)
	if err != nil {
		log.Printf("Error during writing: %+v", err)
	}
	log.Printf("Sent bytes: %d, actual message length: %d", n, len(bytes))
}
