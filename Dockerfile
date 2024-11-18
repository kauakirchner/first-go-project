FROM golang:1.23 AS builder

WORKDIR /app

COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build -o firstgoproject .

FROM golang:1.23-alpine AS runner

RUN adduser -D johndoe

WORKDIR /app

COPY --from=builder /app/firstgoproject /app/firstgoproject
COPY .env ./

RUN chown -R johndoe:johndoe /app
RUN chmod +x /app/firstgoproject

EXPOSE 8080

USER johndoe

CMD ["/app/firstgoproject"]
