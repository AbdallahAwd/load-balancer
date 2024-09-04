package balancer

type Server struct {
	URL     string
	IsAlive bool
}

type LoadBalancer struct {
	Servers       []Server
	CurrentServer int
}

func NewLoadBalancer(servers []Server) *LoadBalancer {
	return &LoadBalancer{
		Servers:       servers,
		CurrentServer: 0,
	}
}

func (lb *LoadBalancer) NextServer() *Server {
	// round robin algorithm
	if len(lb.Servers) == 0 {
		return nil
	}
	lb.CurrentServer = (lb.CurrentServer + 1) % len(lb.Servers)
	return &lb.Servers[lb.CurrentServer]
}
