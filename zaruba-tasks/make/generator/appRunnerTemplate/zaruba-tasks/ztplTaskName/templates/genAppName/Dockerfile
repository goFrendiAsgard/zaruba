FROM golang:1.13 AS builder

RUN mkdir -p /app
WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app


FROM alpine:3.10
COPY --from=builder /app/app /
CMD ["/app"]