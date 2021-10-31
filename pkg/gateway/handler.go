package gateway

import (
	"log"
	"net/http"
)

type gatewayHandler struct {
	clients *grpcClients
}

func NewGatewayHandler(c *grpcClients) *gatewayHandler {
	return &gatewayHandler{clients: c}
}

func errorResponse(err error, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(400)
	w.Write([]byte("bad request"))
}

func (h *gatewayHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("healthy"))
}
