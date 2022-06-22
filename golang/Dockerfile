###############################################################################
## Build stage
###############################################################################
FROM golang:1.18-alpine as build

WORKDIR /app

# Copy manifest
COPY ./go.mod ./

# Install deps
RUN go mod download

# Copy source files
COPY *.go ./

# Build
RUN go build -o /terminal-game-of-life-go



###############################################################################
## Release stage
###############################################################################
FROM debian:buster-slim

# Copy build artefact
COPY --from=build /terminal-game-of-life-go /terminal-game-of-life-go

# Copy entrypoint and make executable
COPY ./docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh

# Run as non-root user
RUN groupadd -r golang-users && useradd -r -g golang-users golang-user
RUN chown golang-user:golang-users /terminal-game-of-life-go
RUN chown golang-user:golang-users /docker-entrypoint.sh
USER golang-user

ENTRYPOINT ["/docker-entrypoint.sh"]
