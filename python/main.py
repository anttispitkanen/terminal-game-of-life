import argparse
import random
import time
from os import system

from blessings import Terminal

from logic import game_of_life_step
from rendering import get_diff_for_rendering, render_diff, render_initial_grid


def create_random_grid(side_length: int) -> list[list[bool]]:
    """
    True represents a live cell,
    False represents a dead cell.
    """
    grid: list[list[bool]] = []
    for y in range(side_length):
        grid.append([])
        for x in range(side_length):
            grid[y].append(random.randint(0, 1) == 1)
    return grid


def parse_args(terminal_height: int, terminal_width: int) -> argparse.Namespace:
    def side_length_type(x):
        """
        The terminal height and width are used to determine the max size of the grid.
        """
        x = int(x)

        # Terminal height as rows. Due to how the grid is printed, we can have max
        # N-1 rows.
        max_height = terminal_height - 1
        # Terminal width as columns, 1 column = 1 character. Since we render "🟪 ",
        # each game cell takes 3 characters. Thus max width is N/3.
        max_width = terminal_width / 3

        max_side_length = int(min(max_height, max_width))

        if x > max_side_length or x < 3:
            raise argparse.ArgumentTypeError(
                f"{x} is not a valid side length, it must be between 3 and {max_side_length} at current terminal window size."
            )

        return x

    parser = argparse.ArgumentParser(
        description="Conway's Game of Life in Terminal in Python"
    )
    parser.add_argument(
        "-s",
        "--side-length",
        type=side_length_type,
        default=20,
        help="Side length of the grid (int).",
    )
    parser.add_argument(
        "-w",
        "--wait-time",
        type=float,
        default=0.4,
        help="Wait time between steps (float)",
    )
    parser.add_argument(
        "-l",
        "--live-emoji",
        type=str,
        default="🟪",
        help="Emoji to use for live cells (str)",
    )
    parser.add_argument(
        "-d",
        "--dead-emoji",
        type=str,
        default="⬜️",
        help="Emoji to use for dead cells (str)",
    )
    return parser.parse_args()


def main():
    try:
        system("clear")
        terminal = Terminal()
        args = parse_args(terminal.height, terminal.width)
        old_grid = create_random_grid(args.side_length)
        new_grid = old_grid.copy()
        render_initial_grid(old_grid, terminal, args.live_emoji, args.dead_emoji)

        while True:
            old_grid = new_grid
            new_grid = game_of_life_step(new_grid)
            diff = get_diff_for_rendering(old_grid, new_grid)
            render_diff(diff, terminal, args.live_emoji, args.dead_emoji)
            time.sleep(args.wait_time)

    except KeyboardInterrupt:
        system("clear")
        print("\nExiting...")


if __name__ == "__main__":
    main()
