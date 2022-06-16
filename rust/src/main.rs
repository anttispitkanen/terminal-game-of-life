use clap::{Arg, Command};
use rand::Rng;
use std::{
    sync::{self, atomic::AtomicBool, Arc},
    thread, time,
};
use termion;
mod logic;
mod rendering;

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

struct Args {
    side_length: i32,
    wait_time: f32,
}

fn parse_args(max_side_length: i32) -> Args {
    let matches = Command::new("Terminal Game of Life – Rust")
        .version("1.0")
        .author("Antti Pitkänen")
        .about("A Rust implementation of the Game of Life in terminal")
        .arg(
            Arg::new("side_length")
                .short('s')
                .long("side_length")
                .value_name("SIDE_LENGTH")
                .help("The side length of the square grid (int)")
                .validator(|f| {
                    let side_length = match f.parse::<i32>() {
                        Ok(side_length) => side_length,
                        Err(_) => {
                            return Err(format!("Invalid side length, must be integer"));
                        }
                    };
                    if side_length > 2 && side_length <= max_side_length {
                        Ok(())
                    } else {
                        Err(format!(
                            "Invalid side length, must be between 3 and {}",
                            max_side_length
                        ))
                    }
                })
                .takes_value(true),
        )
        .arg(
            Arg::new("wait_time")
                .short('w')
                .long("wait_time")
                .value_name("WAIT_TIME")
                .help("Wait time between steps in seconds (float)")
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

fn get_max_side_length(terminal_width: u16, terminal_height: u16) -> i32 {
    // Terminal height as rows. Due to how the grid is printed, we can have max
    // N-1 rows.
    let max_height = terminal_height as i32 - 1;

    // Terminal width as columns, 1 column = 1 character. Since we render "🟪 ",
    // each game cell takes 3 characters. Thus max width is N/3.
    let max_width: i32 = terminal_width as i32 / 3;

    // Return the smaller of the two.
    return std::cmp::min(max_height, max_width);
}

fn main() {
    // Control loop, could this be extracted into a function?
    let running = Arc::new(AtomicBool::new(true));
    let r = running.clone();

    ctrlc::set_handler(move || {
        r.store(false, sync::atomic::Ordering::SeqCst);
    })
    .expect("Error setting SIGINT handler");

    let (terminal_width, terminal_height) =
        termion::terminal_size().expect("Unable to get terminal size");

    let args = parse_args(get_max_side_length(terminal_width, terminal_height));
    rendering::clear_screen();
    let mut old_grid = create_random_grid(args.side_length);
    let mut new_grid = old_grid.clone();
    rendering::print_grid(&old_grid);

    loop {
        old_grid = new_grid.clone();
        new_grid = logic::game_of_life_step(&new_grid);
        let diff = rendering::get_grid_diff_for_rendering(&old_grid, &new_grid);
        rendering::render_diff(&diff);

        thread::sleep(time::Duration::from_secs_f32(args.wait_time));

        if !running.load(sync::atomic::Ordering::SeqCst) {
            break;
        }
    }

    rendering::clear_screen();
    println!("Exiting...");
}
