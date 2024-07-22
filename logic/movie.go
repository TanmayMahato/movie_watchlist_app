package movie

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/TanmayMahato/movie_watchlist_app/database"
	"github.com/TanmayMahato/movie_watchlist_app/models"
	"github.com/lpernett/godotenv"
)

func Getdsn() string {
	err := godotenv.Load("../setup.env")
	if err != nil {
		fmt.Println(err)
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/movieapp"
	return dsn
}

func Insert(w http.ResponseWriter, r *http.Request) error {
	var a1 models.Mvd
	a1.Name = r.FormValue("movie-name")
	a1.Gen = r.FormValue("genre")
	a1.Cat = r.FormValue("industry")
	expval, _ := strconv.ParseInt(r.FormValue("rate"), 10, 64)
	a1.Exp = int(expval)
	database.DBinit(Getdsn())
	mid, erri := database.DBinsert(a1)
	database.DBclose()
	if erri != nil {
		return erri
	}
	movieid := models.Id{
		Movieid: mid,
	}

	tmpl, errt := template.ParseGlob("../../web/add/*.html")
	if errt != nil {
		fmt.Println("Error at main.hdler4.tempParse = ", errt)
		return errt
	}

	if erre := tmpl.ExecuteTemplate(w, "added.html", movieid); erre != nil {
		fmt.Println("Error at main.hdler4.tempExecute = ", erre)
		return erre
	}
	return nil
}

func Show(w http.ResponseWriter, r *http.Request) ([]models.Mvdata, error) {
	var arr []models.Mvdata
	var s2 models.Mvd
	var erra error

	if r.Method == "GET" {
		database.DBinit(Getdsn())
		fmt.Println("hello there123")

		arr, erra = database.DBselect(s2, 1)
		if erra != nil {
			fmt.Println("Error at main.hdler3.stat1")
			return nil, erra
		}
		if err := database.DBclose(); err != nil {
			fmt.Println("Error at main.hdler3.stat2")
		}
		fmt.Println(arr)
		database.DBclose()

	}
	if r.Method == "POST" {
		database.DBinit(Getdsn())
		s2 = models.Mvd{Name: r.FormValue("movie-name"), Gen: r.FormValue("genre"), Cat: r.FormValue("industry")}
		arr, erra = database.DBselect(s2, 0)
		if erra != nil {
			fmt.Println("Error at main.hdler3.stat1") //error
			return nil, erra
		}
		fmt.Println(arr)
		if err := database.DBclose(); err != nil {
			fmt.Println("Error at main.hdler3.stat2")
			return arr, err
		}
		database.DBclose()
	}
	return arr, nil
}

func ShowId(i string) []models.Mvdata {
	var movieid int
	fmt.Println("Check3")
	expval, _ := strconv.ParseInt(i, 10, 64)
	movieid = int(expval)
	fmt.Println(movieid)
	fmt.Println("Check4")
	database.DBinit(Getdsn())
	defer database.DBclose()
	fmt.Println("Check5")
	dat, errm := database.DBidselect(movieid)
	if errm != nil {
		fmt.Println("Error at the movie.ShowId ", errm)
		return nil
	}
	fmt.Println("Check10")
	fmt.Println(dat)
	return dat

}
func UpdateMv(w http.ResponseWriter, r *http.Request) int {
	a := StrtoInt(r.FormValue("hidden"))
	fmt.Println(a)

	data := models.Mvd{Name: r.FormValue("name"), Gen: r.FormValue("genre"), Cat: r.FormValue("industry")}
	database.DBinit(Getdsn())
	no := database.DBupdate(data, a)
	database.DBclose()
	return no

}

func Addtowatched(i string) int {
	a := StrtoInt(i)
	database.DBinit(Getdsn())
	i2 := database.DBdelete(a)
	database.DBclose()

	return i2

}

func StrtoInt(i string) int {
	var movieid int
	fmt.Println("Check3")
	expval, _ := strconv.ParseInt(i, 10, 64)
	movieid = int(expval)
	return movieid

}
