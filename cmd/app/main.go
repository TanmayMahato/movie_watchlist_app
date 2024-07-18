package main

import (
	"log"
	"net/http"
)

func hdler1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/home/index.html")

}

func hdler2(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/add/add.html")
}
func hdler3(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/edit/edit.html")
}
func hdler4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/find/find.html")
}

func hdler5(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/watched/watched.html")
}

func main() {

	http.HandleFunc("/", hdler1)
	http.HandleFunc("/add/", hdler2)
	http.HandleFunc("/edit/", hdler3)
	http.HandleFunc("/find/", hdler4)
	http.HandleFunc("/watched/", hdler5)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
