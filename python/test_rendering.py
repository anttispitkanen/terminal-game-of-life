from rendering import get_diff_for_rendering


def test_get_diff_for_rendering():
    old_grid = [
        [False, False, False],
        [True, True, True],
        [False, False, False],
    ]
    new_grid = [
        [False, True, False],
        [False, True, False],
        [False, True, False],
    ]
    expected_diff = [
        (1, 0, True),
        (0, 1, False),
        (2, 1, False),
        (1, 2, True),
    ]
    assert get_diff_for_rendering(old_grid, new_grid) == expected_diff
