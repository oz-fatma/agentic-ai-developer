# Day 76: multi-stage Docker build demo
# Build: docker build -f day76/Dockerfile -t learn-go-day76 .
# Run from learn/go: docker run --rm learn-go-day76

FROM golang:1.26-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/day76 ./day76

FROM alpine:3.20
RUN adduser -D -g '' app
USER app
COPY --from=builder /bin/day76 /bin/day76
ENTRYPOINT ["/bin/day76"]
