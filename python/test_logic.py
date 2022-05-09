from logic import check_neighbours

test_oscillator = [
    [0, 0, 0],
    [1, 1, 1],
    [0, 0, 0],
]


def test_check_neighbouts():
    # first row
    assert check_neighbours((0, 0), test_oscillator) == 2
    assert check_neighbours((1, 0), test_oscillator) == 3
    assert check_neighbours((2, 0), test_oscillator) == 2

    # second row
    assert check_neighbours((0, 1), test_oscillator) == 1
    assert check_neighbours((1, 1), test_oscillator) == 2
    assert check_neighbours((2, 1), test_oscillator) == 1

    # third row
    assert check_neighbours((0, 2), test_oscillator) == 2
    assert check_neighbours((1, 2), test_oscillator) == 3
    assert check_neighbours((2, 2), test_oscillator) == 2
