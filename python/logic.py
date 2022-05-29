def check_neighbors(coords: tuple[int, int], grid: list[list[bool]]) -> int:
    """
    Check neighbours of a cell, and return the number of alive neighbours.
    """
    alive_neighbors_count = 0
    coord_x = coords[0]
    coord_y = coords[1]
    grid_size_y = len(grid)
    grid_size_x = len(grid[0])

    range_y = range(max(coord_y - 1, 0), min(coord_y + 2, grid_size_y))
    range_x = range(max(coord_x - 1, 0), min(coord_x + 2, grid_size_x))

    for y in range_y:
        for x in range_x:
            # Don't include self in the count
            if grid[y][x] == 1 and (x, y) != coords:
                alive_neighbors_count += 1

    return alive_neighbors_count


def dead_or_alive(original_alive: bool, neigbors_count: int) -> bool:
    """Return the next state of the cell."""
    if original_alive and (neigbors_count == 2 or neigbors_count == 3):
        # remain alive
        return True
    elif not original_alive and (neigbors_count == 3):
        # be born
        return True
    else:
        # die/stay dead
        return False


def game_of_life_step(grid: list[list[bool]]) -> list[list[bool]]:
    """
    Source: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

      1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
      2. Any live cell with two or three live neighbours lives on to the next generation.
      3. Any live cell with more than three live neighbours dies, as if by overpopulation.
      4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

    These rules, which compare the behavior of the automaton to real life, can be condensed into the following:

      1. Any live cell with two or three live neighbours survives.
      2. Any dead cell with three live neighbours becomes a live cell.
      3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
    """
    new_grid: list[list[bool]] = []

    for y, row in enumerate(grid):
        new_row: list[bool] = []
        new_grid.append(new_row)

        for x, val in enumerate(row):
            neighbors = check_neighbors((x, y), grid)
            new_val = dead_or_alive(val, neighbors)
            new_row.append(new_val)

    return new_grid
