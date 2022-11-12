(defproject clojure-tgol "0.1.0-SNAPSHOT"
  :description "Comman line game of life with Clojure"
  :url "https://github.com/anttispitkanen/terminal-game-of-life"
  :dependencies [[org.clojure/clojure "1.11.1"]
                 [org.clojure/tools.cli "1.0.214"]]
  :repl-options {:init-ns clojure-tgol.main}
  :main clojure-tgol.main)
