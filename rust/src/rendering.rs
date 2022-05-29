use super::logic;

// Generic utility
pub fn clear_screen() {
    // Terminal clear sequence, from https://stackoverflow.com/a/66911945
    print!("{esc}c", esc = 27 as char);
}

//
// Initial rendering
//
fn parse_printable_row(row: &logic::Row) -> String {
    let mut row_string = String::new();
    for cell in row {
        if cell == &true {
            row_string.push_str("🟪 ");
        } else {
            row_string.push_str("⬜️ ");
        }
    }
    row_string.push_str("\n");
    return row_string;
}

pub fn print_grid(grid: &logic::Grid) {
    // Clear the screen and position the cursor at (1,1),
    // as inspired by https://hugotunius.se/2019/12/29/efficient-terminal-drawing-in-rust.html
    print!("\x1B[{};{}H", 1, 1);
    for row in grid {
        print!("{}", parse_printable_row(row));
    }
}

//
// Iterative rendering, optimized to only render the diff.
//

// x, y, dead-or-alive
type RenderCoordinate = (usize, usize, bool);

pub fn get_grid_diff_for_rendering(
    old_grid: &logic::Grid,
    new_grid: &logic::Grid,
) -> Vec<RenderCoordinate> {
    let mut diff = Vec::new();

    for (y, row) in old_grid.iter().enumerate() {
        for (x, cell) in row.iter().enumerate() {
            if cell != &new_grid[y][x] {
                diff.push((x, y, new_grid[y][x]));
            }
        }
    }

    return diff;
}

pub fn render_diff(diff: &Vec<RenderCoordinate>, side_length: i32) {
    for (x, y, alive) in diff {
        if *alive {
            println!("\x1B[{};{}H🟪 ", y + 1, (x * 3) + 1);
        } else {
            println!("\x1B[{};{}H⬜️ ", y + 1, (x * 3) + 1);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::{get_grid_diff_for_rendering, RenderCoordinate};

    #[test]
    fn test_get_grid_diff_for_rendering() {
        let old_grid = vec![
            vec![false, false, false],
            vec![true, true, true],
            vec![false, false, false],
        ];
        let new_grid = vec![
            vec![false, true, false],
            vec![false, true, false],
            vec![false, true, false],
        ];

        let diff = get_grid_diff_for_rendering(&old_grid, &new_grid);

        let expected_diff: Vec<RenderCoordinate> =
            vec![(1, 0, true), (0, 1, false), (2, 1, false), (1, 2, true)];

        assert_eq!(diff, expected_diff);
    }
}
