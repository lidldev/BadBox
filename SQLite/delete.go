package main

import (
	"fmt"
	_"github.com/mattn/go-sqlite3"
	"log"
	"database/sql"
)

func checkErr5(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

func main() {
	db,err:=sql.Open("sqlite3","filmler.db")
	defer db.Close()
	checkErr5(err)
	
	silme,err:=db.Prepare("DELETE FROM oyuncular WHERE id=?")
	checkErr5(err)
	
	res,err:=silme.Exec(4)
	checkErr5(err)
	
	affect,err:=res.RowsAffected()
	checkErr5(err)
	fmt.Println("Silinen Satır Sayısı:",affect)
}