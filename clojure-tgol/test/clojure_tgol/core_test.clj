(ns clojure-tgol.core-test
  (:require [clojure.test :refer :all]
            [clojure-tgol.core :refer :all]))

(def test-grid
  [[false false false]
   [true true true]
   [false false false]])

(deftest check-neighbors-test
  (is (= (check-neighbors [0 0] test-grid) 2))
  (is (= (check-neighbors [1 0] test-grid) 3))
  (is (= (check-neighbors [2 0] test-grid) 2))
  (is (= (check-neighbors [0 1] test-grid) 1))
  (is (= (check-neighbors [1 1] test-grid) 2))
  (is (= (check-neighbors [2 1] test-grid) 1))
  (is (= (check-neighbors [0 2] test-grid) 2))
  (is (= (check-neighbors [1 2] test-grid) 3))
  (is (= (check-neighbors [2 2] test-grid) 2)))

(deftest game-of-life-step-test
  (let [expected-output-grid
        [[false, true, false]
         [false, true, false]
         [false, true, false]]
        actual-output-grid (game-of-life-step test-grid)]
    ; One iteration of game of life on the test grid should result in the expecte
    ; output grid.
    (is (= actual-output-grid expected-output-grid))
    ; The next iteration on the output grid should result in the original grid.
    (is (= (game-of-life-step actual-output-grid) test-grid))))

