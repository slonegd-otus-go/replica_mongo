FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ./client/client ./client/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/client/client .

ENTRYPOINT ["./client"]