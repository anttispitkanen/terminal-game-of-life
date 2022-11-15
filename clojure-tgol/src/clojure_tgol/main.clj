(ns clojure-tgol.main
  (:gen-class)
  (:require [clojure-tgol.core :refer [game-of-life-step]]
            [clojure-tgol.rendering :refer [get-diff-for-rendering render-diff render-initial-grid]]
            [clojure.tools.cli :refer [parse-opts]]
            [clojure.java.shell :refer [sh]]
            [clojure.string :as str]))

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

(defn get-terminal-dimensions
  []
  (as-> (sh "/bin/sh" "-c" "stty size < /dev/tty") t
    (:out t)
    (clojure.string/trim t)
    (clojure.string/split t #" ")
    {:rows (Integer/parseInt (first t)) :cols (Integer/parseInt (second t))}))


(defn get-max-side-length
  []
  (as-> (get-terminal-dimensions) t
    (let [;; Since each cell is 3 printed characters wide, we need to divide the
          ;; terminal width by 3
          max-widh (/ (:cols t) 3)
          max-height (:rows t)]
      (min max-widh max-height))))

;; This is a function instead of a constant so that it doesn't need to be evaluated
;; at uberjar compilation time. That was a problem in the docker build, apparently
;; because the /dev/tty device doesn't exist at that time, and that's needed for
;; getting the terminal dimensions used here.
(defn get-cli-options []
  [["-w" "--wait-time WAIT_TIME_IN_SECONDS" "Wait time between steps in seconds"
    :default 0.4
    :parse-fn #(Float/parseFloat %)
    :validate [#(< 0.01 % 3) "Must be a number between 0.01 and 3"]]
   ["-s" "--side-length SIDE_LENGTH" (format "Side length (between 3 and %d)" (get-max-side-length))
    :default 20
    :parse-fn #(Integer/parseInt %)
    :validate [#(< 4 % (get-max-side-length)) (format "Must be a number between 3 and %d" (get-max-side-length))]]
   ["-h" "--help"]])

;; TODO:
;; - Handle keyboard interrupt
;;   - This turned out not to be so trivial, maybe looking at it later again

(defn -main [& args]
  (let [opts (parse-opts args (get-cli-options))
        options (:options opts)
        summary (:summary opts)
        errors (:errors opts)]
    ;; Help requested => prints and exits
    (if (:help options) ((println summary) (System/exit 0)) nil)
    ;;(if errors (println "ei mittää :D") (System/exit 0)))))
    (if errors ((println errors) (System/exit 1)) nil)
    ;; No help requested => runs the game
    (let [grid (create-random-grid (:side-length options))]
      (clear-terminal)
      (render-initial-grid grid)
      (loop [current-grid grid]
        (let [new-grid (game-of-life-step current-grid)]
          (render-diff (get-diff-for-rendering current-grid new-grid))
          (Thread/sleep (* (:wait-time options) 1000))
          (recur new-grid))))))
