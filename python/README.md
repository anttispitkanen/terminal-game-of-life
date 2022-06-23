# Terminal Game of Life – Python

You need python 3. If your alias points to python 2, replace `python` with `python3`, or fix the situation with something like pyenv.

## Setup

Setup the virtual env

```bash
python -m venv venv
source venv/bin/activate
```

Install dependencies

```bash
pip install -r requirements.txt
```

## Running

Run the program, `-h|--help` will show you the arguments and what they do

```bash
python main.py -h
```

## Running in Docker

```bash
# Build
docker build -t terminal-game-of-life-python .
# Run
docker run -it terminal-game-of-life-python [args]
```

Note that you need the `-it` flag to pass Docker the tty, otherwise you'll get an error as the program can't read it.

Also note the small delay in the program starting. This is explained in the [root README](../README.md).

## Testing

Run tests with

```bash
pytest .
```
