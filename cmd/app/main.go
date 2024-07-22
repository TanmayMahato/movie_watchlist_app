package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/TanmayMahato/movie_watchlist_app/database"
	movie "github.com/TanmayMahato/movie_watchlist_app/logic"
	"github.com/TanmayMahato/movie_watchlist_app/models"
	"github.com/lpernett/godotenv"
)

func hdler1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/home/index.html")

}

func hdler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "../../web/add/add.html")
		return
	}

	err := movie.Insert(w, r)
	if err != nil {
		fmt.Println("Eror at hdler2 ", err)
		return
	}

}

func hdler3(w http.ResponseWriter, r *http.Request) {

	arr, errshow := movie.Show(w, r)

	if errshow != nil {
		fmt.Println("Error at movie.Show = ", errshow)
		return
	}
	tmpl, errt := template.ParseGlob("../../web/show/*.html")
	if errt != nil {
		fmt.Println("Error at main.hdler3.tempParse = ", errt)
		return
	}
	if erre := tmpl.ExecuteTemplate(w, "show.html", arr); erre != nil {
		fmt.Println("Error at main.hdler3.tempExecute = ", erre)
		return
	}
}

func hdler4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Check1")
	idholder = r.FormValue("hidden")
	http.Get("http://localhost:8080/editfind")
	fmt.Println("Check")

	tmpl, errt := template.ParseGlob("../../web/edit/*.html")
	if errt != nil {
		fmt.Println("Error at main.hdler4.tempParse = ", errt)
		return
	}
	dat := data
	fmt.Println(dat)
	fmt.Println("Check12")
	if erre := tmpl.ExecuteTemplate(w, "edit.html", dat); erre != nil {
		fmt.Println("Error at main.hdler4.tempExecute = ", erre)
		return
	}

}

var idholder string
var data []models.Mvdata

func hdler4o5(w http.ResponseWriter, r *http.Request) {
	var num []int
	if r.Method == "GET" {
		fmt.Println("Check2")
		data = movie.ShowId(idholder)
		fmt.Println("Check11", data)
	}
	if r.Method == "POST" {
		err := movie.UpdateMv(w, r)
		if err != 1 {
			fmt.Println("Something Went Wrong ..")
			return
		}
		if err == 1 {
			fmt.Println("Updated successfully ")
			num = append(num, 1)
		}

		tmpl, errt := template.ParseGlob("../../web/find/*.html")
		if errt != nil {
			fmt.Println("Error at main.hdler4.tempParse = ", errt)
			return
		}

		if erre := tmpl.ExecuteTemplate(w, "editfind.html", num); erre != nil {
			fmt.Println("Error at main.hdler4.tempExecute = ", erre)
			return
		}

	}

}

func hdler5(w http.ResponseWriter, r *http.Request) {

	database.DBinit(DSN)
	data, errd := database.DBwatchedselect()
	database.DBclose()
	if errd != nil {
		fmt.Println("Error at hdler5", errd)
		return
	}

	tmpl, errt := template.ParseGlob("../../web/watched/*.html")
	if errt != nil {
		fmt.Println("Error at main.hdler4.tempParse = ", data)
		return
	}

	if erre := tmpl.ExecuteTemplate(w, "watched.html", data); erre != nil {
		fmt.Println("Error at main.hdler4.tempExecute = ", erre)
		return
	}

}

func hdler6(w http.ResponseWriter, r *http.Request) {
	var d []int

	if r.Method == "POST" {
		val := r.FormValue("hidden")
		i2 := movie.Addtowatched(val)
		if i2 != 1 {
			fmt.Println("Error at Adding to the watched hdler6")
			return
		} else {
			d = append(d, 1)
		}

		tmpl, errt := template.ParseGlob("../../web/watched/*.html")
		if errt != nil {
			fmt.Println("Error at main.hdler4.tempParse = ", errt)
			return
		}

		if erre := tmpl.ExecuteTemplate(w, "addedtowatched.html", d); erre != nil {
			fmt.Println("Error at main.hdler4.tempExecute = ", erre)
			return
		}

	}

}

var DSN string

func Getdsn() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/movieapp"
	DSN = dsn

}

func main() {
	err := godotenv.Load("../../setup.env")
	if err != nil {
		fmt.Println(err)
	}
	Getdsn()

	http.HandleFunc("/", hdler1)
	http.HandleFunc("/add", hdler2)
	http.HandleFunc("/show", hdler3)
	http.HandleFunc("/edit", hdler4)
	http.HandleFunc("/editfind", hdler4o5)
	http.HandleFunc("/watched", hdler5)
	http.HandleFunc("/addwatched", hdler6)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
