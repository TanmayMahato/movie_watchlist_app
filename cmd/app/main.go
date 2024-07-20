package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/TanmayMahato/movie_watchlist_app/database"
	"github.com/TanmayMahato/movie_watchlist_app/models"
)

func hdler1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/home/index.html")

}

func hdler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "../../web/add/add.html")
		return
	}
	// con, _ := io.ReadAll(r.Body)
	// res := string(con)
	// fmt.Println(res)
	var a1 models.Mvd
	a1.Name = r.FormValue("movie-name")
	a1.Gen = r.FormValue("genre")
	a1.Cat = r.FormValue("industry")
	expval, _ := strconv.ParseInt(r.FormValue("rate"), 64, 10)
	a1.Exp = int(expval)
	dsn := "root:2003@tcp(localhost:3306)/movieapp"
	database.DBinit(dsn)
	mid := database.DBinsert(a1)
	database.DBclose()

	movieid := models.Id{
		Movieid: mid,
	}

	tmpl, _ := template.ParseFiles("templates/com_addconf.html")
	_ = tmpl.Execute(w, movieid)

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
	http.HandleFunc("/add", hdler2)
	http.HandleFunc("/edit/", hdler3)
	http.HandleFunc("/find/", hdler4)
	http.HandleFunc("/watched/", hdler5)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
