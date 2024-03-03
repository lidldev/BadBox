package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Company struct{
	Cliste []Sirket
}

type Sirket struct{
	ID int
	Isim string
	KurulusTarihi int
	DevamEdiyorMu bool
	
}

func main() {
	s1:=Sirket{}
	s1.ID=1
	s1.Isim="Yazılım Merkezi"
	s1.KurulusTarihi=2013
	s1.DevamEdiyorMu=true
	
	sirket1byte,err:=json.Marshal(s1)
	if err!=nil{
		log.Fatal(err)
	}
	
	fmt.Println(string(sirket1byte))
	
	s2:=Sirket{55,"Microsoft",1975,true}
	sirket2byte,_:=json.Marshal(s2)
	
	fmt.Println(string(sirket2byte))
	
	dosya,err:=os.Create("sirketler.json")
	if err!=nil{
		log.Fatal(err)
	}
	
	c1:=Company{}
	
	c1.Cliste=append(c1.Cliste,Sirket{ID:1,Isim:"Kızılay",KurulusTarihi:1868,DevamEdiyorMu:true})
	c1.Cliste=append(c1.Cliste,Sirket{ID:2,Isim:"Microsoft",KurulusTarihi:1975,DevamEdiyorMu:true})
	
	jsonWriter:=io.Writer(dosya)
	enc:=json.NewEncoder(jsonWriter)
	json.MarshalIndent(c1," ","\t")
	enc.Encode(c1)
	
	
}