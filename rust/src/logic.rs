pub type Row = Vec<bool>;
pub type Grid = Vec<Row>;

fn check_neighbors(x_coord: usize, y_coord: usize, grid: &Grid) -> usize {
    let mut alive_neighbors: usize = 0;
    let grid_size_y: usize = grid.len();
    let grid_size_x: usize = grid[0].len();

    let y_start = if y_coord == 0 { 0 } else { y_coord - 1 };
    let y_end = if y_coord == grid_size_y - 1 {
        y_coord
    } else {
        y_coord + 1
    };

    let x_start = if x_coord == 0 { 0 } else { x_coord - 1 };
    let x_end = if x_coord == grid_size_x - 1 {
        x_coord
    } else {
        x_coord + 1
    };

    for y in y_start..=y_end {
        for x in x_start..=x_end {
            if grid[y][x] == true && !(x == x_coord && y == y_coord) {
                alive_neighbors += 1;
            }
        }
    }

    return alive_neighbors;
}

/// Take a grid and run one iteration of the game of life.
/// Return a new grid.
pub fn game_of_life_step(grid: &Grid) -> Grid {
    /*
      Source: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

      1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
      2. Any live cell with two or three live neighbours lives on to the next generation.
      3. Any live cell with more than three live neighbours dies, as if by overpopulation.
      4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

      These rules, which compare the behavior of the automaton to real life, can be condensed into the following:

      1. Any live cell with two or three live neighbours survives.
      2. Any dead cell with three live neighbours becomes a live cell.
      3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
    */
    let mut new_grid: Grid = Vec::new();

    let mut y: usize = 0; // store y index

    for row in grid {
        let mut x: usize = 0; // store x index
        let mut new_row: Row = Vec::new();

        for cell in row {
            let alive_neighbors_count = check_neighbors(x, y, grid);
            let original_alive = cell == &true;

            if original_alive && (alive_neighbors_count == 2 || alive_neighbors_count == 3) {
                // Remain alive
                new_row.push(true);
            } else if !original_alive && alive_neighbors_count == 3 {
                // Be born
                new_row.push(true);
            } else {
                // Die or stay dead
                new_row.push(false);
            }

            x += 1; // increment x index
        }
        new_grid.push(new_row);

        y += 1; // increment y index
    }

    return new_grid;
}

#[cfg(test)]
mod tests {
    use super::*;

    fn create_test_grid() -> Grid {
        return vec![
            vec![false, false, false],
            vec![true, true, true],
            vec![false, false, false],
        ];
    }

    #[test]
    fn test_check_neighbors() {
        let test_grid = create_test_grid();

        // First row
        assert_eq!(2, check_neighbors(0, 0, &test_grid));
        assert_eq!(3, check_neighbors(1, 0, &test_grid));
        assert_eq!(2, check_neighbors(2, 0, &test_grid));

        // Second row
        assert_eq!(1, check_neighbors(0, 1, &test_grid));
        assert_eq!(2, check_neighbors(1, 1, &test_grid));
        assert_eq!(1, check_neighbors(2, 1, &test_grid));

        // Third row
        assert_eq!(2, check_neighbors(0, 2, &test_grid));
        assert_eq!(3, check_neighbors(1, 2, &test_grid));
        assert_eq!(2, check_neighbors(2, 2, &test_grid));
    }

    #[test]
    fn test_game_of_life_step() {
        let test_grid = create_test_grid();

        let expected_grid = vec![
            vec![false, true, false],
            vec![false, true, false],
            vec![false, true, false],
        ];

        let result_grid = game_of_life_step(&test_grid);

        assert_eq!(expected_grid, result_grid);

        let second_result_grid = game_of_life_step(&result_grid);

        assert_eq!(test_grid, second_result_grid);
    }
}
