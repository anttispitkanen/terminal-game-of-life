# Terminal Game of Life – Clojure

Note the name of the directory, a Leiningen app cannot be named "clojure", so it had to be suffixed with "-tgol" 🤷

## Prerequisites

- [Clojure](https://clojure.org/guides/install_clojure)
  - Note its prerequisite of Java, follow the link above
- [Leiningen](https://leiningen.org/)

## Running

Note that `lein run` and the actual args (like `-w 0.1 -s 30`) need to be separated by two dashes `--`. Without it the args are considered for `lein run` itself, and not the app.

```bash
lein run [-- [args]]
```

You probably also want to set up a [REPL](https://clojure.org/guides/repl/introduction) for the development, but that's outside the scope of this document.

## Building

```bash
lein uberjar
```

After which you can run the jar, provided that you have Java installed:

```bash
java -jar ./target/clojure-tgol-0.1.0-SANPSHOT-standalone.jar [args]
```

## Running in Docker

```bash
# Build
docker build -t terminal-game-of-life-clojure .
# Run
docker run -it terminal-game-of-life-clojure [args]
```

Note that you need the `-it` flag to pass Docker the tty, otherwise you'll get an error as the program can't read it.

## Testing

```bash
lein test
```
