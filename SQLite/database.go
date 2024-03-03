package main

import (
	"fmt"
	_"github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
)

func checkErr4(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

func main() {
	db,err:=sql.Open("sqlite3","filmler.db")
	defer db.Close()
	checkErr4(err)
	
	guncelle,err:=db.Prepare("UPDATE oyuncular SET ad=? WHERE id=?")
	checkErr4(err)
	res,err:=guncelle.Exec("Ali",2)
	fmt.Println(res)
	
}