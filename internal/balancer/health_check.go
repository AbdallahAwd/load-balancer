package balancer

import (
	"net/http"
	"time"
)

func (s *Server) ServerHealthCheck() {
	timeout := time.Second * 2
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(s.URL)
	if err != nil || resp.StatusCode != 200 {
		s.IsAlive = false
	} else {
		s.IsAlive = true
	}
}

func (lb *LoadBalancer) PerfornHealthCheck() {
	for i := 0; i < len(lb.Servers); i++ {
		lb.Servers[i].ServerHealthCheck()
	}
}
