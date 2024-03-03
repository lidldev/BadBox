package main

//SMTP -- Simple Mail Transfer Protocol

import (
	"fmt"
	"log"
	"net/smtp"
)

//SMTP mail göndermek
//POP IMAP mail almak

func main() {

	// Gönderici Bilgileri
	gonderen := "golangahks@gmail.com"
	sifre := "Udemy444"

	// Alıcı Bilgiler
	alici := []string{
		"musleraneuercasillas@gmail.com",
		}

		// STMP Ayarları
		smtpHost := "smtp.gmail.com"
		smtpPort := "587"

		// Mesaj.
		baslik := "Affan"
		icerik := "Naber Lan"

		mesaj := []byte(baslik + icerik)

		//Authentication İşlemleri

		auth := smtp.PlainAuth("", gonderen, sifre, smtpHost)

		// Email Gönderme

		for i:=0;i<100;i++{
			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, gonderen, alici, mesaj)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("E-Posta başarıyla gönderildi.")

}