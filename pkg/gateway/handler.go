package gateway

import (
	"net/http"
)

type gatewayHandler struct {
	clients *grpcClients
}

func NewGatewayHandler(c *grpcClients) *gatewayHandler {
	return &gatewayHandler{clients: c}
}

func (h *gatewayHandler) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("healthy"))
}
