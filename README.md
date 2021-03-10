# grpc-load-balancer-application

Distributed-system for keeping track of users expenses.

The system has two Go services that communicate with each other via gRPC (worker and loadBalancer).

## How to start project

***

*Add .env file in the root of the project*
```
# Network settings
PORT=3001

# DB setup
DATABASE_URL=postgres://postgres:example@localhost:5155/postgres?sslmode=disable
REGISTRY_URL=redis://localhost:6300

# JWT secret
JWT_SECRET=secret
```

***

*Run docker-compose*
> cd dev\
> docker-compose up

***

*Install dependencies*
> go mod tidy

***

*Run migration*
> make migrateup

***

*Start load balancer application*
> go run cmd/loadBalancer/*.go

***

*Start workers*
> PORT=9001 go run cmd/worker/main.go\
> PORT=9002 go run cmd/worker/main.go
