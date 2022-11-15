(ns clojure-tgol.core)

(defn get-neighbor-values
  "Given a coordinate and a grid, return the values of the relevan neighbors in a single vector."
  [coords grid]
  (let [coord-x (first coords)
        coord-y (second coords)
        grid-size-y (count grid)
        grid-size-x (count (first grid))
        range-y (range (max (- coord-y 1) 0) (min (+ coord-y 2) grid-size-y))
        range-x (range (max (- coord-x 1) 0) (min (+ coord-x 2) grid-size-x))]
    (for [y range-y]
      (for [x range-x]
        (if (and (= true (nth (nth grid y) x)) (not (= [x y] coords)))
          1 0)))))

(defn check-neighbors
  "Given a coordinate and a grid, count the number of alive neighbors."
  [coords grid]
  (->> (get-neighbor-values coords grid)
       (apply concat)
       (reduce + 0)))


(defn dead-or-alive
  "Return the next state of the cell."
  [original-alive neigbors-count]
  (if (and original-alive (or (= neigbors-count 2) (= neigbors-count 3)))
    ;; remain alive
    true
    (if (and (not original-alive) (= neigbors-count 3))
      ;; be born
      true
      ;; die/stay dead
      false)))


(defn game-of-life-step
  ;;     Source: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
  ;;       1. Any live cell with fewer than two live neighbours dies, as if by underpopulation.
  ;;       2. Any live cell with two or three live neighbours lives on to the next generation.
  ;;       3. Any live cell with more than three live neighbours dies, as if by overpopulation.
  ;;       4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
  ;;     These rules, which compare the behavior of the automaton to real life, can be condensed into the following:
  ;;       1. Any live cell with two or three live neighbours survives.
  ;;       2. Any dead cell with three live neighbours becomes a live cell.
  ;;       3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
  "Takes in a grid, runs a single iteration of game of life on it, and returns the resulting grid."
  [grid]
  (let [grid-size-y (count grid)
        grid-size-x (count (first grid))]
    (for [y (range grid-size-y)]
      (for [x (range grid-size-x)]
        (dead-or-alive (nth (nth grid y) x) (check-neighbors [x y] grid))))))

