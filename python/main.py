import random
import time
from os import system
from typing import Literal, cast

from blessings import Terminal

from logic import game_of_life_step


def create_random_grid(side_length: int) -> list[list[Literal[1, 0]]]:
    """
    1s represent live cells,
    0s represent dead cells.
    """
    grid: list[list[Literal[1, 0]]] = []
    for y in range(side_length):
        grid.append([])
        for x in range(side_length):
            grid[y].append(cast(Literal[1, 0], random.randint(0, 1)))
    return grid


def parse_printable_row(row: list[Literal[1, 0]], terminal: Terminal) -> str:
    parsed_row = ""
    for val in row:
        if val == 1:
            parsed_row += terminal.black("🟪 ")
        else:
            parsed_row += terminal.white("⬜️ ")
    return parsed_row


def print_grid(grid: list[list[Literal[1, 0]]], terminal: Terminal):
    for i, row in enumerate(grid):
        with terminal.location(x=0, y=i):
            print(parse_printable_row(row, terminal))


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
        grid = create_random_grid(args.side_length)

        while True:
            print_grid(grid, terminal)
            grid = game_of_life_step(grid)
            time.sleep(args.wait_time)

    except KeyboardInterrupt:
        system("clear")
        print("\nExiting...")


if __name__ == "__main__":
    main()
