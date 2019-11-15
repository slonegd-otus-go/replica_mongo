FROM golang:1.13
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o client1 ./client/main.go
ENTRYPOINT ["./client1"]


