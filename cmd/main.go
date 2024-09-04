package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AbdallahAwd/laod-balancer/config"
	"github.com/AbdallahAwd/laod-balancer/internal/balancer"
	"github.com/AbdallahAwd/laod-balancer/internal/handlers"
	ser "github.com/AbdallahAwd/laod-balancer/internal/server"
)

func main() {
	conf, err := config.LoadConfig(".env")
	if err != nil {
		panic(err)
	}
	servers := []balancer.Server{}
	for _, v := range conf.Backends {
		servers = append(servers, balancer.Server{URL: v})
	}
	lb := balancer.NewLoadBalancer(servers)
	lb.PerfornHealthCheck()
	handler := handlers.NewLoadBalancerHandler(lb)

	http.HandleFunc("/", handler.ServeHTTP)

	server := http.Server{
		Addr:    conf.Port,
		Handler: handler,
	}
	go func() {
		if err := ser.Start(&server); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()
	shutdownGracefully(&server)
}

func shutdownGracefully(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server stopped gracefully")
}
