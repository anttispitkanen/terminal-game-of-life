# terminal-game-of-life

Building [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in terminal in different languages. This is my pet project to try out new languages.

![Video demo](/demo.gif)

## Spec

This is how every individual implementation should work regardless of language. It must:

- Be launched with a single CLI command
- Support CLI arguments for
  - `-s|--side-length` = grid size = side length of a square grid, default to 20
    - This must be limited to not be larger than can fit in the current terminal window
  - `-w|--wait-time` = the time between each increment in the game, default to 400ms
  - `-h|--help` print help message
- According to given arguments, print the game in the terminal window and update it in place as seen in the demo above
- Handle keyboard interrupt, gracefully clearing the game off the terminal
- Be able to run with Docker too

### Design philosophy

- The implementation between languages should be as similar as possible, while respecting each language's conventions over unity
- The logic and data model should be completely separate from how the visuals are rendered, it should be possible to utilize the same core for a different rendering method
- The rendering should be optimized to not require full redraw of the whole grid for each tick
- Simple unit tests should be enough to validate the core logic functions

### Program structure

- `Main` module with functions

  - `Main` – handles command flow, the game loop, clearing the terminal, and handling keyboard interruption
  - `Random grid generation` – parametrized side length
    - `true` represents an alive cell, `false` a dead cell
  - `Argument parsing` – encapsulates reading the CLI args and exposing those to the control flow
    - Here we need to make sure the side length argument is not smaller than 3, and not larger than what can fit on the current terminal window

- `Logic` module with functions

  - `Check neighbors` – count the number of alive neighbors for a given cell coordinate according to the Game of Life rules
  - `Dead or alive` – based on the cell's previous state and neighbor count, return the cell's next state
  - `Game of life step` – take in a grid and **return a new grid** with the Game of Life rules applied to it, by running the check neighbors function for each cell, and then the dead or alive function

- `Rendering` module with functions

  - `Clear screen` – clearing the screen before and after a game
  - For initial rendering:
    - `Get row for rendering` – the representation of alive cells as "🟪 " and dead cells as "⬜️ "
    - `Render initial grid` – prints the grid in place by calling the print row function for each row
  - For iterative rendering:
    - `Get diff for rendering` – loops over the previous grid and the current grid, and returns a list of the cells that need to be rendered again, in the format of a coordinate pair `(x, y)` and the new state for that cell (`true`/`false`)
    - `Render diff` – takes the output from `get diff for rendering` as input, and prints the cells that need to be updated, overwriting the previous state of those cells

### Tests

#### Logic tests

The `logic` module should have tests for

- The `check neighbors` function

  If given an input grid like this

  ```json
  [
    [false, false, false],
    [true, true, true],
    [false, false, false]
  ]
  ```

  it should return, for the corresponding coordinates, values (per cell)

  ```json
  [
    [2, 3, 2],
    [1, 2, 1],
    [2, 3, 2]
  ]
  ```

- The `game of life step` function

  If given an input grid like this

  ```json
  [
    [false, false, false],
    [true, true, true],
    [false, false, false]
  ]
  ```

  the output should look like this

  ```json
  [
    [false, true, false],
    [false, true, false],
    [false, true, false]
  ]
  ```

  and when giving that as an input, it should return the first grid again.

#### Rendering tests

The `rendering` module should have tests for

- The `get diff for rendering` function

  If given an old grid of

  ```json
  [
    [false, false, false],
    [true, true, true],
    [false, false, false]
  ]
  ```

  and a new grid of

  ```json
  [
    [false, true, false],
    [false, true, false],
    [false, true, false]
  ]
  ```

  the resulting diff should be

  ```json
  [
    [1, 0, true],
    [0, 1, false],
    [2, 1, false],
    [1, 2, true]
  ]
  ```

### Docker

We want a multi-stage build where the final resulting image is as lean as possible and running as non-root.

#### Note about terminal in Docker

Due to a bug in Docker, in order to correctly get the dimensions of the active host terminal, we need to `sleep 1` before running the script. Here's a repo dedicated to researching and explaining why that is: https://github.com/anttispitkanen/rust-terminal-dimensions-in-docker
