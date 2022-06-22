# Terminal Game of Life - Go

You need golang 1.18.

## Setup

No external dependencies needed.

## Running

Run the program without building, `-h|--help` will show you the arguments and what they do

```bash
go run .
```

Or build the program into a binary and run that

```bash
go build . # => builds the executable binary "terminal-game-of-life"
./terminal-game-of-life # => runs the binary
```

## Running in Docker

```bash
# Build
docker build -t terminal-game-of-life-go .
# Run
docker run -it terminal-game-of-life-go [args]
```

Note that you need the `-it` flag to pass Docker the tty, otherwise you'll get an error as the program can't read it.

Also note the small delay in the program starting. This is explained in the [root README](../README.md).

## Testing

Run tests with

```bash
go test .
```
