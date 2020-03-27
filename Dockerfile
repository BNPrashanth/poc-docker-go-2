FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add git
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-two-server ./cmd/go-two-server/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/ /app/
EXPOSE 8082
EXPOSE 8082
CMD ["./go-two-server"]