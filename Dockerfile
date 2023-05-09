FROM golang:1.20.4-bullseye AS builder

WORKDIR /app

COPY . .

RUN go build -o tj

FROM alpine:3.17.3

COPY --from=builder /app/tj /usr/bin/tj

CMD ["tj"]
