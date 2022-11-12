(ns clojure-tgol.main
  (:gen-class)
  (:require [clojure-tgol.core :refer [game-of-life-step]]
            [clojure-tgol.rendering :refer [get-diff-for-rendering render-diff render-initial-grid]]
            [clojure.tools.cli :refer [parse-opts]]))

(defn clear-terminal []
  (print (str (char 27) "c")))

(defn get-true-or-false
  []
  (if (= (rand-int 2) 1) true false))

(defn create-random-row
  [side-length]
  (vec (take side-length (repeatedly get-true-or-false))))

(defn create-random-grid
  [side-length]
  (vec (take side-length (repeatedly #(create-random-row side-length)))))

(def cli-options
  [["-w" "--wait-time WAIT_TIME_IN_SECONDS" "Wait time between steps in seconds"
    :default 0.4
    :parse-fn #(Float/parseFloat %)
    :validate [#(< 0.01 % 3) "Must be a number between 0.01 and 3"]]])

(defn -main [& args]
  (as-> (parse-opts args cli-options) opts
    (:wait-time (:options opts))
    (let [grid (create-random-grid 20)]
      (clear-terminal)
      (render-initial-grid grid)
      (loop [current-grid grid]
        (Thread/sleep (* opts 1000)) ;; <- This is the wait time
        (let [new-grid (game-of-life-step current-grid)]
          (render-diff (get-diff-for-rendering current-grid new-grid))
          (recur new-grid))))))
