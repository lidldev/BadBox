package main

import (
	"database/sql"
	"fmt"
	"log"
	_"github.com/mattn/go-sqlite3"
)

func tickErr(err error){
	if err !=nil{
		log.Fatal(err)
	}
}

func main() {
	db,err:=sql.Open("sqlite3","filmler.db")
	defer db.Close()
	tickErr(err)
	
	satirlar,err:=db.Query("SELECT * FROM filmler")
	tickErr(err)
	
	var id int
	var filmAd string
	var filmKategori string
	var yonetmenAd string
	var yapimyili int
	var butce int
	
	for satirlar.Next(){ 
		err = satirlar.Scan(&id, &filmAd, &filmKategori, &yonetmenAd, &yapimyili, &butce)
		tickErr(err)
		fmt.Println("ID:",id,"Filmin Adı:",filmAd,"Filmin Kategorisi:",filmKategori,"Yönetmen Adı:",yonetmenAd,"Yapım Yılı:",yapimyili,"Bütçe",butce)
	}
	
}