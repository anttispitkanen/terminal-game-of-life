# Terminal Game of Life – Rust

You need Rust >= 1.61.

## Running

The debug run is convenient while developing. Note that you need to separate the arguments with `--` as they are not args to `cargo run` but to the program itself.

```bash
cargo run [-- [args]]
```

## Building

```bash
# Debug build
cargo build
# Release build
cargo build --release
```

After which you can run the binary with

```bash
# Debug build
./target/debug/rust [args]
# Release build
./target/release/rust [args]
```

## Running in Docker

```bash
# Build
docker build -t terminal-game-of-life-rust .
# Run
docker run -it terminal-game-of-life-rust [args]
```

Note that you need the `-it` flag to pass Docker the tty, otherwise you'll get an error as the program can't read it.

Also note the small delay in the program starting. This is explained in the [root README](../README.md).

## Testing

```bash
cargo test
```
