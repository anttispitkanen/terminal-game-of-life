def check_neighbours(coords: tuple[int, int], grid: list[list[int]]) -> int:
    """
    Check neighbours of a cell, and return the number of alive neighbours.
    """
    alive_neighbours_count = 0
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
                alive_neighbours_count += 1

    return alive_neighbours_count


def game_of_life_step(grid: list[list[int]]):
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
    new_grid: list[list[int]] = []

    for y, row in enumerate(grid):
        new_row: list[int] = []
        new_grid.append(new_row)

        for x, val in enumerate(row):
            original_alive = val == 1
            neighbours = check_neighbours((x, y), grid)

            if original_alive and (neighbours == 2 or neighbours == 3):
                # remain alive
                new_grid[y].append(1)
            elif not original_alive and (neighbours == 3):
                # be born
                new_grid[y].append(1)
            else:
                # die/stay dead
                new_grid[y].append(0)

    return new_grid
