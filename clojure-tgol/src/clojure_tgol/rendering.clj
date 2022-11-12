(ns clojure-tgol.rendering)

;;
;; Initial rendering
;;

(defn parse-printable-row
  "Parse a row of the grid into a printable string."
  [row]
  (as-> row r
    (map (fn [cell] (if cell "🟪 " "⬜️ ")) r)
    (apply str r)
    (str r "\n")))

(defn render-initial-grid
  "Render the initial grid."
  [grid]
  (->> grid
       (map parse-printable-row)
       (apply str)
       print)
  (flush))

;;
;; Iterative rendering, focused on only rendering the diff
;;
(defn get-diff-for-rendering
  "Get the diff between the previous grid and current grid for rendering."
  [old-grid new-grid]
  (let [range-y (range (count old-grid))
        range-x (range (count (first old-grid)))]
    (filter (fn [value] (not (nil? value)))
            (apply concat
                   (for [y range-y]
                     (for [x range-x]
                       (if (not (= (nth (nth old-grid y) x) (nth (nth new-grid y) x)))
                         [x y (nth (nth new-grid y) x)] nil)))))))

(defn render-diff
  "Render the diff between the previous grid and current grid. Just prints, doesn't return anything."
  [diff]
  (doseq [value diff]
    (let [x (first value)
          y (second value)
          alive (nth value 2)]
      (if alive
        (print (format "\033[%s;%sH🟪 " (+ y 1) (+ (* x 3) 1)))
        (print (format "\033[%s;%sH⬜️ " (+ y 1) (+ (* x 3) 1))))))
  (flush))
