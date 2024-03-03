package main

import (
	"database/sql"
	"log"
	_"github.com/mattn/go-sqlite3"
)

func checkErr3(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

func main() {
	db,err:=sql.Open("sqlite3","filmler.db")
	defer db.Close()
	checkErr3(err)
	
	oyuncuAd:=[]string{"Sener Sen", "Ali"}
	oyuncuId:=[]int {1,3}
	oyuncuYas:=[]int{74, 68}

	for i:=0;i<len(oyuncuId);i++{
		veriekle,err:=db.Prepare("INSERT INTO oyuncular (id, ad , yas) VALUES (? , ? , ?) ")
		checkErr3(err)
		veriekle.Exec (oyuncuId[i], oyuncuAd[i], oyuncuYas [i])
	}	
}