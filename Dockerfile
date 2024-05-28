# Step 1: Build the Go binary
FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN go build -o dist/connect4 cmd/main.go

# Step 2: Create a small image to run the Go binary
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/dist/connect4 .

# Command to run the executable
CMD ["./connect4"]