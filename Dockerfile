FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app
COPY --from=builder /app/replica_mongo .

ENTRYPOINT ["./replica_mongo"]


