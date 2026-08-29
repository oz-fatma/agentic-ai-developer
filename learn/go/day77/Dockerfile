FROM golang:1.26-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app ./day77

FROM alpine:3.20
EXPOSE 8080
COPY --from=builder /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
