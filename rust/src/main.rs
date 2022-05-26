use rand::Rng;
use std::{thread, time};
mod logic;

fn get_random_cell() -> bool {
    let mut rng = rand::thread_rng();
    return rng.gen_range(0..2) == 1;
}

fn create_random_grid(side_length: i32) -> logic::Grid {
    let mut grid = Vec::new();
    for _ in 0..side_length {
        let mut row = Vec::new();
        for _ in 0..side_length {
            row.push(get_random_cell());
        }
        grid.push(row);
    }
    return grid;
}

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

fn print_grid(grid: &logic::Grid) {
    // Clear the screen and position the cursor at (1,1),
    // as inspired by https://hugotunius.se/2019/12/29/efficient-terminal-drawing-in-rust.html
    print!("\x1B[{};{}H", 1, 1);
    for row in grid {
        print!("{}", parse_printable_row(row));
    }
}

fn clear_screen() {
    // Terminal clear sequence, from https://stackoverflow.com/a/66911945
    print!("{esc}c", esc = 27 as char);
}

fn main() {
    clear_screen();
    let mut grid = create_random_grid(20);

    loop {
        print_grid(&grid);
        grid = logic::game_of_life_step(&grid);
        thread::sleep(time::Duration::from_millis(400));
    }
}
