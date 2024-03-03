package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

type Students struct {
	XMLName        xml.Name `xml:"Ogrenciler"`
	OgrenciListesi []Ogr    `xml:"ogrenci"`
}

type Ogr struct {
	XMLName      xml.Name `xml:"ogrenci"`
	OgrenciNo    int      `xml:"ogrNo"`
	OgrenciAd    string   `xml:"ogrAd"`
	OgrenciSoyad string   `xml:"ogrSoy"`
	OgrenciNot   int      `xml:"ogrNot"`
}

func main() {
	dosya, err := os.Open("ogrenci.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer dosya.Close()

	arr, err := io.ReadAll(dosya)
	if err != nil {
		log.Fatal(err)
	}

	Ogrenci := &Students{}
	xml.Unmarshal(arr, Ogrenci)

	for _, degerler := range Ogrenci.OgrenciListesi {
		fmt.Println("Öğrenci Adı:\n", degerler.OgrenciAd)
		fmt.Println("Öğrenci Soyadı:\n",degerler.OgrenciSoyad)
		fmt.Println("Öğrenci Notu:\n",degerler.OgrenciNot)
		fmt.Println("Öğrenci No'su\n",degerler.OgrenciNo)
	}
}