FROM golang:1.13.8-alpine3.11 AS builder
WORKDIR /golang-wiki
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o wiki main.go

FROM scratch
COPY --from=builder /golang-wiki/wiki /golang-wiki/wiki
WORKDIR /golang-wiki
COPY src/views src/views
WORKDIR tmp
WORKDIR /golang-wiki

ENTRYPOINT ["/golang-wiki/wiki"]

EXPOSE 8080
