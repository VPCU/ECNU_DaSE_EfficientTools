FROM golang:1.20-alpine
COPY server.go .
CMD go run server.go
