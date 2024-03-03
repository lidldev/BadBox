package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {

	mesaj := gomail.NewMessage()

	mesaj.SetHeader("From", "ahmetaffan123@gmail.com")
	mesaj.SetHeader("To", "ahmetaffan123@gmail.com")
	mesaj.SetHeader("Subject", "Önemli!!!")

	mesaj.Attach("udemy.jpg")

	mesaj.SetBody("text/html", "<p>Bu bir paragraf</p><h2 style=\"text-align:center\">Bunu Gördün mü?</h2>")

	dia := gomail.NewDialer("smtp.gmail.com", 587, "ahmetaffan123@gmail.com", "Ahmet1718")

	err := dia.DialAndSend(mesaj)
	if err != nil {
		log.Fatal(err)
	}

}
