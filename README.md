<h1>Структура директории</h1>

/pr2
│
├── docker-compose.yml
├── docs/
│   └── openapi.yaml
│
├── api_gateway/
│   ├── cmd/
│   │   └── gateway/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   ├── cors.go
│   │   │   └── request_id.go
│   │   ├── proxy/
│   │   │   ├── users.go
│   │   │   └── orders.go
│   │   └── server/
│   │       └── http.go
│   ├── Dockerfile
│   └── go.mod
│
├── service_users/
│   ├── cmd/
│   │   └── users/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── db/
│   │   │   └── repository.go
│   │   ├── handlers/
│   │   │   ├── auth.go
│   │   │   └── profile.go
│   │   ├── models/
│   │   │   └── user.go
│   │   └── service/
│   │       └── user_service.go
│   ├── Dockerfile
│   └── go.mod
│
├── service_orders/
│   ├── cmd/
│   │   └── orders/
│   │       └── main.go
│   ├── internal/
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── db/
│   │   │   └── repository.go
│   │   ├── handlers/
│   │   │   ├── create.go
│   │   │   ├── get.go
│   │   │   └── update.go
│   │   ├── models/
│   │   │   └── order.go
│   │   └── service/
│   │       └── order_service.go
│   ├── Dockerfile
│   └── go.mod
│
└── pkg/
    ├── jwt/
    │   └── jwt.go
    ├── logger/
    │   └── logger.go
    └── tracing/
        └── tracing.go
