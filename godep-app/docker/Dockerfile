# Dockerfile

FROM golang:1.16 AS builder

WORKDIR /app

COPY . .

RUN go get -u github.com/tools/godep

RUN godep restore

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
RUN chmod +x main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
