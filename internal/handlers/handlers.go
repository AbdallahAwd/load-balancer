package handlers

import (
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

	if nextServer == nil || !nextServer.IsAlive {
		h.ServeHTTP(w, r)
		nextServer.ServerHealthCheck()
		return
	}

	log.Printf("Sending request to %s", nextServer.URL)

	http.Redirect(w, r, nextServer.URL, http.StatusFound)
}
