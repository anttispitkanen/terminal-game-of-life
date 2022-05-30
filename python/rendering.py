from blessings import Terminal


#
# Initial rendering.
#
def parse_printable_row(row: list[bool], terminal: Terminal) -> str:
    parsed_row = ""
    for val in row:
        if val:
            parsed_row += terminal.black("🟪 ")
        else:
            parsed_row += terminal.white("⬜️ ")
    return parsed_row


def render_initial_grid(grid: list[list[bool]], terminal: Terminal):
    for i, row in enumerate(grid):
        with terminal.location(x=0, y=i):
            print(parse_printable_row(row, terminal))


#
# Iterative rendering, optimized to only render the diff.
#
def get_diff_for_rendering(
    old_grid: list[list[bool]], new_grid: list[list[bool]]
) -> list[tuple[int, int, bool]]:
    diff = []
    for y, row in enumerate(old_grid):
        for x, val in enumerate(row):
            if val != new_grid[y][x]:
                diff.append((x, y, new_grid[y][x]))
    return diff


def render_diff(diff: list[tuple[int, int, bool]], terminal: Terminal):
    for x, y, is_alive in diff:
        with terminal.location(x=x * 3, y=y):
            print("🟪 " if is_alive else "⬜️ ")
