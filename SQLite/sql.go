package main

import (
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"log"
)

func checkErr1(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

func main() {
	db,err:=sql.Open("sqlite3","filmler.db")
	defer db.Close()
	checkErr1(err)
	db.Exec("CREATE TABLE notlar(id INT PRIMARY KEY,noturumu TEXT,isim TEXT)")
	
}