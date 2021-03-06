from logic import check_neighbors, game_of_life_step

test_grid = [
    [False, False, False],
    [True, True, True],
    [False, False, False],
]


def test_check_neighbors():
    # first row
    assert check_neighbors((0, 0), test_grid) == 2
    assert check_neighbors((1, 0), test_grid) == 3
    assert check_neighbors((2, 0), test_grid) == 2

    # second row
    assert check_neighbors((0, 1), test_grid) == 1
    assert check_neighbors((1, 1), test_grid) == 2
    assert check_neighbors((2, 1), test_grid) == 1

    # third row
    assert check_neighbors((0, 2), test_grid) == 2
    assert check_neighbors((1, 2), test_grid) == 3
    assert check_neighbors((2, 2), test_grid) == 2


def test_game_of_life_step():
    # First iteration
    expected_grid = [
        [False, True, False],
        [False, True, False],
        [False, True, False],
    ]
    output_1 = game_of_life_step(test_grid)
    assert output_1 == expected_grid

    # Second iteration
    output_2 = game_of_life_step(output_1)
    assert output_2 == test_grid
