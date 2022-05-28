use clap::{Arg, Command};
use rand::Rng;
use std::{
    sync::{self, atomic::AtomicBool, Arc},
    thread, time,
};
mod logic;

fn random_cell_value() -> bool {
    let mut rng = rand::thread_rng();
    return rng.gen_range(0..2) == 1;
}

fn create_random_grid(side_length: i32) -> logic::Grid {
    let mut grid = Vec::new();
    for _ in 0..side_length {
        let mut row = Vec::new();
        for _ in 0..side_length {
            row.push(random_cell_value());
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

struct Args {
    side_length: i32,
    wait_time: f32,
}

fn parse_args() -> Args {
    let matches = Command::new("Terminal Game of Life – Rust")
        .version("1.0")
        .author("Antti Pitkänen")
        .about("A Rust implementation of the Game of Life in terminal")
        .arg(
            Arg::new("side_length")
                .short('s')
                .long("side_length")
                .value_name("SIDE_LENGTH")
                .help("The side length of the square grid")
                .takes_value(true),
        )
        .arg(
            Arg::new("wait_time")
                .short('w')
                .long("wait_time")
                .value_name("WAIT_TIME")
                .help("Wait time between steps in seconds")
                .takes_value(true),
        )
        .get_matches();

    let side_length = matches
        .value_of("side_length")
        .unwrap_or("20")
        .parse::<i32>()
        .unwrap();
    let wait_time = matches
        .value_of("wait_time")
        .unwrap_or("0.4")
        .parse::<f32>()
        .unwrap();

    return Args {
        side_length,
        wait_time,
    };
}

fn main() {
    // Control loop, could this be extracted into a function?
    let running = Arc::new(AtomicBool::new(true));
    let r = running.clone();

    ctrlc::set_handler(move || {
        r.store(false, sync::atomic::Ordering::SeqCst);
    })
    .expect("Error setting SIGINT handler");

    let args = parse_args();
    clear_screen();
    let mut grid = create_random_grid(args.side_length);

    // Could this be done the other way around? I.e. "if false then break"
    while running.load(sync::atomic::Ordering::SeqCst) {
        print_grid(&grid);
        grid = logic::game_of_life_step(&grid);
        thread::sleep(time::Duration::from_secs_f32(args.wait_time));
    }

    clear_screen();
    println!("Exiting...");
}
