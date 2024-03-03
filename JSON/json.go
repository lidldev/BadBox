package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Students struct{
	Ogrenciler []Ogrenci `json:"ogrenciler"`
}

type Ogrenci struct{
	OgrenciNo int `json:"ogrNo"`
	OgrenciAd string `json:"ogrAd"`
	OgrenciSoyad string `json:"ogrSoyad"`
	OgrenciNot int `json:"ogrNot"`
}

func main() {
	dosya,err:=os.Open("ogrenciler.json")
	if err!=nil{
		log.Fatal(err)
	}
	
	arr,err:=ioutil.ReadAll(dosya)
	if err!=nil{
		log.Fatal(err)
	}
	
	var Ogr Students
	
	json.Unmarshal(arr,&Ogr)
	
	fmt.Println(string(arr))
	
	for i:=0;i<len(Ogr.Ogrenciler);i++{
		fmt.Println("Öğrenci No'su:",Ogr.Ogrenciler[i].OgrenciNo)
		fmt.Println("Öğrenci Adı:",Ogr.Ogrenciler[i].OgrenciAd)
		fmt.Println("Öğrenci Soyadı:",Ogr.Ogrenciler[i].OgrenciSoyad)
		fmt.Println("Öğrenci Notu:",Ogr.Ogrenciler[i].OgrenciNot)
	}
}