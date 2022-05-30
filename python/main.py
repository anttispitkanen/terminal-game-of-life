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


def parse_args():
    import argparse

    parser = argparse.ArgumentParser(
        description="Conway's Game of Life in Terminal in Python"
    )
    parser.add_argument(
        "-s", "--side-length", type=int, default=20, help="Side length of the grid"
    )
    parser.add_argument(
        "-w", "--wait-time", type=float, default=0.4, help="Wait time between steps"
    )
    return parser.parse_args()


def main():
    try:
        system("clear")
        terminal = Terminal()
        args = parse_args()
        old_grid = create_random_grid(args.side_length)
        new_grid = old_grid.copy()
        render_initial_grid(old_grid, terminal)

        while True:
            old_grid = new_grid
            new_grid = game_of_life_step(new_grid)
            diff = get_diff_for_rendering(old_grid, new_grid)
            render_diff(diff, terminal)
            time.sleep(args.wait_time)

    except KeyboardInterrupt:
        system("clear")
        print("\nExiting...")


if __name__ == "__main__":
    main()
