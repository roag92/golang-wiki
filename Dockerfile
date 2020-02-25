FROM golang:1.13.8-alpine3.11 AS base
WORKDIR /golang-wiki
COPY . .
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
RUN GO111MODULE=on go get -v golang.org/x/tools/gopls@latest 2>&1
USER $USERNAME

FROM base AS builder
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o wiki main.go

FROM scratch
COPY --from=builder /golang-wiki/wiki /golang-wiki/wiki
WORKDIR /golang-wiki
COPY src/views src/views
WORKDIR tmp
WORKDIR /golang-wiki
ENTRYPOINT ["/golang-wiki/wiki"]
