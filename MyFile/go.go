package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("Kitap Arama Uygulamasına Hoş Geldiniz!")
	fmt.Println("Aramak istediğiniz kitabın adını yazın, Enter tuşuna basın. Çıkmak için 'exit' yazabilirsiniz.")

	for {
		fmt.Print("Kitap adı veya Enter tuşuna basın: ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Uygulama kapatılıyor...")
			break
		}

		if input != "" {
			searchTerm := strings.ReplaceAll(input, " ", "+")
			searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", searchTerm)
			openBrowser(searchURL)
		}
	}
}

func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
		case "windows":
			cmd = exec.Command("cmd", "/c", "start", url)
			default:
				cmd = exec.Command("xdg-open", url)
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Tarayıcı açılamadı:", err)
	}
}
