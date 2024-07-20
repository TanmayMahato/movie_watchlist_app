package ad1

import (
	"log"
	"net/http"
)

type Mvdata struct {
	Name string
	Gen  string
	Cat  string
	Exp  int
}

func adhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	// datas := mvdata{Name: r.FormValue("name"), Gen: r.FormValue("genre"), Cat: r.FormValue("industry")}
	// fmt.Println(datas)
}
func main() {

	http.HandleFunc("/add_data", adhandler)
	log.Fatal(http.ListenAndServe(":8040", nil))
}
