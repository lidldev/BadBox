package main

import (
	"fmt"
	_"github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
)

func checkErr6(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

func main() {
	db,err:=sql.Open("sqlite3","filmler.db")
	defer db.Close()
	checkErr6(err)
	
	guncelle,err:=db.Prepare("UPDATE filmler SET butce=?")
	checkErr6(err)
	
	tx,err:=db.Begin()
	checkErr6(err)
	
	_,txerr:=tx.Stmt(guncelle).Exec(500)
	if txerr !=nil{
		fmt.Println("Geri Alınıyor...ROLLBACK")
		tx.Rollback()
	}else{
		fmt.Println("Commit Edildi")
		tx.Commit()
	}
	
}