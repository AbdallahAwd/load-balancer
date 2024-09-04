package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AbdallahAwd/laod-balancer/internal/balancer"
)

type LoadBalancerHandler struct {
	LoadBalancer *balancer.LoadBalancer
}

func NewLoadBalancerHandler(lb *balancer.LoadBalancer) *LoadBalancerHandler {
	return &LoadBalancerHandler{
		LoadBalancer: lb,
	}
}

func (h *LoadBalancerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	nextServer := h.LoadBalancer.NextServer()
	log.Print("In the Server / Handler ", nextServer.URL, " Alive ", nextServer.IsAlive)
	if nextServer == nil || !nextServer.IsAlive {
		json.NewEncoder(w).Encode(map[string]string{"message": "No servers available"})
		return
	}

	log.Printf("Sending request to %s", nextServer.URL)

	http.Redirect(w, r, nextServer.URL, http.StatusFound)
}
