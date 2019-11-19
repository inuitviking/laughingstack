(ns hahaha
  [:require [clojure.core.async :refer [go go-loop <!!]]])

(defn laugh []
  "ha")

(defn join-laughs [laughs]
  (clojure.string/join laughs))

(defn laugh-by-composition []
  (let [f (comp println join-laughs)]
    (f [(laugh) (laugh) (laugh)])))

(defn laugh-forever []
  (repeat (laugh)))

(defn laugh-using-take []
  (println (join-laughs (take 3 (laugh-forever)))))

(defn laugh-using-map []
  (doall (map (fn [l] (print l)) (take 3 (laugh-forever)))))

(defn laugh-asyncronously []
  (go-loop [i 0]
    (go (print (laugh)))
    (if (< i 2)
      (recur (inc i)))))

(defn laugh-async-block []
  (<!! (laugh-asyncronously)))

(defn laugh-in-a-lot-of-ways []
  (for [f [laugh-by-composition laugh-using-take laugh-async-block laugh-using-map]]
    (f)))

(defn -main []
  (doall (laugh-in-a-lot-of-ways)))
