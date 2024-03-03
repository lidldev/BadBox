package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	k_adi := "AAE"
	k_sifre := "admin"
	girishakki := 3
	fmt.Println("*******LOGIN*******")
	for true {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Kullanıcı Adınızı Giriniz:")
		scanner.Scan()
		giris_kadi := scanner.Text()
		fmt.Print("Şifrenizi Giriniz:")
		scanner.Scan()
		giris_ksifre := scanner.Text()
		if k_adi != giris_kadi && k_sifre != giris_ksifre {
			fmt.Println("Kullanıcı Adı ve Şifreniz Hatalı")
			girishakki--
		} else if k_adi == giris_kadi && k_sifre != giris_ksifre {
			fmt.Println("Şifreniz Hatalı")
			girishakki--
		} else if k_adi != giris_kadi && k_sifre == giris_ksifre {
			fmt.Println("Kullanıcı Adınız Hatalı")
			girishakki--
		} else {
			fmt.Println("Giriş Başarılı")
			break
		}
		if girishakki == 0 {
			fmt.Println("Giriş Hakkınız Bitmiştir")
			break
		}
	}

}
