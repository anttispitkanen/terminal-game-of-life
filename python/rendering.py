import unicodedata

from blessings import Terminal


def _get_character_width(char: str) -> int:
    """
    As emojis are unicode characters and thus have a len(emoji) == 1, we need to check
    the unicode.east_asian_width(emoji) == "W" to consider the character as having a
    render width of 2 slots.
    """
    return 2 if unicodedata.east_asian_width(char[0]) in "W" else len(char)


def _get_rendered_character(char: str) -> str:
    """
    All characters need to be in total 3 slots wide, with whitespace added as padding.
    As emojis occupy 2 slots, we need to add just one whitespace after them.
    """
    ONE_SPACE = " "
    TWO_SPACES = "  "
    return char + TWO_SPACES if _get_character_width(char) == 1 else char + ONE_SPACE


#
# Initial rendering.
#
def parse_printable_row(
    row: list[bool], terminal: Terminal, alive_cell: str, dead_cell: str
) -> str:
    ac = _get_rendered_character(alive_cell)
    dc = _get_rendered_character(dead_cell)

    parsed_row = ""
    for val in row:
        if val:
            parsed_row += terminal.black(ac)
        else:
            parsed_row += terminal.white(dc)
    return parsed_row


def render_initial_grid(
    grid: list[list[bool]],
    terminal: Terminal,
    alive_cell: str,
    dead_cell: str,
):
    for i, row in enumerate(grid):
        with terminal.location(x=0, y=i):
            print(parse_printable_row(row, terminal, alive_cell, dead_cell))


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


def render_diff(
    diff: list[tuple[int, int, bool]],
    terminal: Terminal,
    alive_cell: str,
    dead_cell: str,
):
    ac = _get_rendered_character(alive_cell)
    dc = _get_rendered_character(dead_cell)

    for x, y, is_alive in diff:
        with terminal.location(x=x * 3, y=y):
            print(ac if is_alive else dc)
