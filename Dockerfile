FROM cosmtrek/air

WORKDIR /app
COPY main.go main.go
COPY go.* .