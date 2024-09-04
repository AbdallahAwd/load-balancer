# Simple Load Balancer

## Overview

This project is a simple load balancer implemented in Go, designed to distribute incoming HTTP requests across multiple backend servers. It uses basic round-robin load balancing and includes features for health checks and graceful shutdown.

## Features

- **Load Balancing**: Distributes incoming requests using a round-robin algorithm.
- **Health Checks**: Monitors the status of backend servers to ensure availability.
- **Graceful Shutdown**: Handles interrupt signals to shut down the server gracefully.
- **Logging**: Captures logs for debugging and monitoring purposes.

## Project Structure

```
 ├── cmd/
 │ └── main.go
 ├── internal/
 │ ├── balancer/
 │ │ └── balancer.go
 │ └── handlers/
 │ └── handlers.go
 ├── .gitignore
 ├── go.mod
 └── README.md
```


- **cmd/main.go**: Entry point of the application. Initializes the load balancer and starts the HTTP server.
- **internal/balancer/balancer.go**: Contains the load balancer logic, including backend server management and load balancing algorithms.
- **internal/handlers/handlers.go**: Defines HTTP request handlers that interact with the load balancer to route requests to backend servers.

## Setup and Installation

1. **Clone the Repository**

   ```sh
   git clone https://github.com/AbdallahAwd/load-balancer.git
   cd load-balancer
2. **Add .env**

   ```sh
   PORT=localhost:8888
   BACKENDS=server1,server2
3. **Run Project**

   ```sh
   go run .\cmd\
