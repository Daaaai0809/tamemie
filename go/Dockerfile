FROM golang:1.23.4-alpine3.21 AS dev

ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

RUN apk add --no-cache curl

# install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin

EXPOSE 8000
CMD ["go", "run", "cmd/server/main.go"]

FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o tamemie cmd/server/main.go

FROM alpine AS prod

WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/tamemie /app/tamemie
EXPOSE 8000
CMD ["/app/tamemie"]
