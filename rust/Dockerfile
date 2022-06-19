###############################################################################
## Build stage
###############################################################################
FROM rust:1.61-slim-buster as build

# create new empty sehll project
RUN USER=root cargo new --bin terminal-game-of-life-rust
WORKDIR /terminal-game-of-life-rust

# copy manifests
COPY ./Cargo.* ./

# build and cache deps
RUN cargo build --release
RUN rm src/*.rs

# copy source tree
COPY ./src ./src

# build for release
RUN rm ./target/release/deps/*
RUN cargo build --release

###############################################################################
## Release stage
###############################################################################
FROM debian:buster-slim

# copy build artefact
COPY --from=build /terminal-game-of-life-rust/target/release/rust .

# copy entrypoint script and make it executable
COPY ./docker-entrypoint.sh .
RUN chmod +x ./docker-entrypoint.sh

# Run as non-root
RUN groupadd -r rust-users && useradd -r -g rust-users rust-user
RUN chown rust-user:rust-users /rust
RUN chown rust-user:rust-users /docker-entrypoint.sh
USER rust-user

ENTRYPOINT ["./docker-entrypoint.sh"]
