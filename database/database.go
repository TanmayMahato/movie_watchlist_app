package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TanmayMahato/movie_watchlist_app/models"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// initializes the db
func DBinit(dsn string) {
	db, _ = sql.Open("mysql", dsn)
	fmt.Println("Connected to the db!")
}

// db close
func DBclose() error {
	if db != nil {
		db.Close()
		fmt.Println(" db closed!")
		return nil
	}
	return nil
}

// func to add a movie in watchlist table .. returns the last inserted mvid no
func DBinsert(s1 models.Mvd) (int, error) {
	var ins *sql.Stmt
	var err2 error
	ins, err2 = db.Prepare("INSERT INTO `movieapp`.`watchlist`(`name` ,`gen` ,`cat` ,`exp`) VALUES(?,?,?,?);")
	if err2 != nil {
		fmt.Println("Error at dbinsert while prepare ", err2)
		return 5, err2
	}
	defer ins.Close()

	res, err := ins.Exec(s1.Name, s1.Gen, s1.Cat, s1.Exp)
	if err != nil {
		fmt.Println("Error at dbinsert while exec ", err)
		return 5, err
	}
	ls, errl := res.LastInsertId()
	if errl != nil {
		fmt.Println("Error at database.DBInsert while lastinsertid ", errl)
		return int(ls), errl
	}

	return int(ls), nil
}

// id = 0 for not all and 1 for select all along with a Mvd struct as input
// returns the watchlist table rows in a slice also the row gets added to watched table with the trigger in the db
func DBselect(list models.Mvd, id int) ([]models.Mvdata, error) {
	var response *sql.Rows
	var errs error
	if id == 1 {
		response, errs = db.Query("SELECT * FROM `movieapp`.`watchlist`;")
	} else if id == 0 && list.Name == "" && list.Gen == "" && list.Cat == "" {
		response, errs = db.Query("SELECT * FROM `movieapp`.`watchlist`;")
	} else if id == 0 && list.Name != "" && list.Gen == "" && list.Cat == "" {
		name := list.Name + "%"
		response, errs = db.Query("SELECT * FROM `movieapp`.`watchlist` where `name` LIKE ?;", name)
	} else if id == 0 && list.Name == "" && list.Gen != "" && list.Cat == "" {
		response, errs = db.Query("SELECT * FROM `movieapp`.`watchlist` where `gen`= ?;", list.Gen)
	} else if id == 0 && list.Name == "" && list.Gen == "" && list.Cat != "" {
		response, errs = db.Query("SELECT * FROM `movieapp`.`watchlist` where `cat`= ?;", list.Cat)
	} else {
		response, errs = db.Query("SELECT * FROM `movieapp`.`watchlist`;")
	}
	if errs != nil {
		return nil, errs
	}
	defer response.Close()

	var Arrdata []models.Mvdata

	for response.Next() {
		var ad models.Mvdata

		if err := response.Scan(&ad.Mvid, &ad.Name, &ad.Gen, &ad.Cat, &ad.Exp); err != nil {
			return Arrdata, err
		}
		Arrdata = append(Arrdata, ad)

	}
	if err := response.Err(); err != nil {
		return Arrdata, err
	}

	return Arrdata, nil

}

// returns the watched table rows in a slice
func DBwatchedselect() ([]models.Mvwdata, error) {

	response, _ := db.Query("SELECT * FROM `movieapp`.`watched`;")
	var Arrdata []models.Mvwdata
	defer response.Close()

	for response.Next() {
		var ad models.Mvwdata
		err := response.Scan(&ad.Mvid, &ad.Name, &ad.Gen, &ad.Cat, &ad.Rate)
		if err != nil {
			return nil, err
		}
		Arrdata = append(Arrdata, ad)
	}

	return Arrdata, nil

}

// Select with the id
func DBidselect(i int) ([]models.Mvdata, error) {
	fmt.Println("Check6 ", i)
	res := db.QueryRow("SELECT * FROM `movieapp`.`watchlist` where `mvid` = ? ;", i)
	// if err != nil {
	// 	fmt.Println("Error at database.DBidselect1", err)
	// 	return nil, err
	// }
	fmt.Println("Check7")
	var Arrdata []models.Mvdata
	// defer res.Close()

	fmt.Println("Check8 - one time only")
	var ad models.Mvdata
	err := res.Scan(&ad.Mvid, &ad.Name, &ad.Gen, &ad.Cat, &ad.Exp)
	if err != nil {
		return nil, err
	}
	Arrdata = append(Arrdata, ad)
	fmt.Println(Arrdata)

	fmt.Println("Check9")
	return Arrdata, nil

}

// deletes the row with the given id
func DBdelete(id int) int {
	var del *sql.Stmt
	var err3 error
	del, err3 = db.Prepare("Delete from `movieapp`.`watchlist` where (`mvid`=?);")
	if err3 != nil {
		log.Fatal(err3)
	}
	defer del.Close()

	res, err := del.Exec(id)

	if err != nil {
		panic(err.Error())
	}

	raffected, _ := res.RowsAffected()
	if raffected == 1 {
		fmt.Println("Succesfully deleted the row with id -> ", id)
		return 1 // Success or true
	}
	return 0 // false or failed

}

// update the name gen cat
func DBupdate(s4 models.Mvd, id int) int {
	var upt *sql.Stmt
	var err4 error
	upt, err4 = db.Prepare("UPDATE `movieapp`.`watchlist` SET `name` = ?, `gen` = ?, `cat`= ? , `exp`=? WHERE `mvid`=? ;")
	if err4 != nil {
		log.Fatal(err4)
	}
	defer upt.Close()

	res, err := upt.Exec(s4.Name, s4.Gen, s4.Cat, s4.Exp, id)

	if err != nil {
		panic(err.Error())
	}
	if xy, _ := res.RowsAffected(); xy == 1 {
		fmt.Println("successfully updated the data row")
		return 1 // true or success
	}
	return 0 // false or failed

}
