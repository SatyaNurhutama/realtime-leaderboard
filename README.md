# Real-Time Leaderboard with Go, Redis, and WebSocket

This project demonstrates a real-time leaderboard using Go, Redis, WebSocket, and Gin. Users can submit scores, retrieve the top scores, and see updates in real-time.

## Features

- Submit and store scores in Redis sorted sets
- Fetch top N leaderboard entries
- Broadcast score updates via WebSocket

## Architecture
```
├── cmd/
│   └── main.go                 # Entry point
├── pkg/
│   └── redis_client/           # Redis client initialization
├── internal/
│   └── leaderboard/            # Leaderboard domain logic
│       ├── redis_store.go      # Redis operations
│       ├── service.go          # Business logic
│       ├── handler.go          # HTTP handlers
│       └── websocket.go        # WebSocket broadcasting
```
