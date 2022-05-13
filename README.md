# terminal-game-of-life

Building [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) in terminal in different languages. Or maybe just Python if I never get back to this 😆

![Video demo](/demo.gif)

## Spec

This is how every individual implementation should work regardless of language. It must:

- Be launched with a single CLI command
- Support CLI arguments for
  - `-s|--side-length` = grid size = side length of a square grid, default to 20
  - `-w|--wait-time` = the time between each increment in the game, default to 400ms
  - `-h|--help` print help message
- According to given arguments, print the game in the terminal window and update it in place as seen in the demo above
- Handle keyboard interrupt, gracefully clearing the game off the terminal

### Program structure

- `Main` module with functions

  - `Main` – handles command flow, the game loop, clearing the terminal, and handling keyboard interruption
  - `Random grid generation` – parametrized side length
    - `1` represents an alive cell, `0` a dead cell
  - `Printable row parsing` – the representation of alive cells as "🟪 " and dead cells as "⬜️ "
  - `Print grid` – prints the grid in place (erasing the previously printed grid) by calling the print row function for each row
  - `Argument parsing` – encapsulates reading the CLI args and exposing those to the control flow

- `Logic` module with functions
  - `Check neighbors` – count the number of alive neighbors for a given cell coordinate according to the Game of Life rules
  - `Dead or alive` – based on the cell's previous state and neighbor count, return the cell's next state
  - `Game of life step` – take in a grid and **return a new grid** with the Game of Life rules applied to it, by running the check neighbors function for each cell, and then the dead or alive function

#### Tests

The `logic` module should have tests for

- The `check neighbors` function

  if given an input grid like this

  ```json
  [
    [0, 0, 0],
    [1, 1, 1],
    [0, 0, 0]
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
    [0, 0, 0],
    [1, 1, 1],
    [0, 0, 0]
  ]
  ```

  the output should look like this

  ```json
  [
    [0, 1, 0],
    [0, 1, 0],
    [0, 1, 0]
  ]
  ```

  and when giving that as an input, it should return the first grid again.
