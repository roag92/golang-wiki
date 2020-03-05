FROM golang:1.13.8-alpine3.11 AS base
WORKDIR /go/src/golang-wiki
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go mod download 2>&1
EXPOSE 8080

FROM base AS dev
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000
RUN adduser $USERNAME -s /bin/sh -D -u $USER_UID $USER_GID; \
    mkdir -p /etc/sudoers.d; \
    echo "$USERNAME ALL=(root) NOPASSWD:ALL" > /etc/sudoers.d/$USERNAME; \
    chmod 0440 /etc/sudoers.d/$USERNAME
RUN apk add -q --update --progress --no-cache git sudo bash curl
RUN GO111MODULE=on go get -v golang.org/x/tools/gopls@latest \
    github.com/ramya-rao-a/go-outline 2>&1
RUN chmod +rx /go/pkg/ -R
USER $USERNAME
RUN sudo chown $USERNAME:$USERNAME /go/src/golang-wiki -R

FROM base AS builder
USER root
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /go/src/golang-wiki/bin/golang-wiki main.go

FROM scratch
COPY --from=builder /go/src/golang-wiki/bin/golang-wiki /usr/local/golang-wiki/bin/golang-wiki
WORKDIR /usr/local/golang-wiki
COPY .env templates tmp/.gitkeep ./
ENTRYPOINT ["/usr/local/golang-wiki/bin/golang-wiki"]
