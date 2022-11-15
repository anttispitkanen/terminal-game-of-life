(ns clojure-tgol.rendering-test
  "Testing the rendering logic"
  (:require [clojure.test :refer [deftest is]]
            [clojure-tgol.rendering :refer [get-diff-for-rendering]]))

(def test-initial-grid [[false false false]
                        [true true true]
                        [false false false]])

(def test-new-grid [[false true false]
                    [false true false]
                    [false true false]])

(deftest get-diff-for-rendering-test
  (is (= (get-diff-for-rendering test-initial-grid test-new-grid)
         [[1 0 true] [0 1 false] [2 1 false] [1 2 true]])))