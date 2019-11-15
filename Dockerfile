FROM golang:1.13
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build
ENTRYPOINT ["./replica_mongo"]


