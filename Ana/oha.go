package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Event struct {
	ID      int
	Title   string
	Date    string
	Details string
}

var events = make(map[int]Event)
var lastEventID = 0

func main() {
	fmt.Println("Etkinlik Yöneticisine Hoş Geldiniz!")

	for {
		fmt.Println("Seçenekleri belirtin:")
		fmt.Println("1. Etkinlik Listesi")
		fmt.Println("2. Etkinlik Ekle")
		fmt.Println("3. Etkinlik Düzenle")
		fmt.Println("4. Etkinlik Sil")
		fmt.Println("5. Çıkış")

		var choice int
		fmt.Print("Seçiminizi girin: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			listEvents()
			case 2:
				addEvent()
				case 3:
					editEvent()
					case 4:
						deleteEvent()
						case 5:
							fmt.Println("Çıkıyor...")
							os.Exit(0)
							default:
								fmt.Println("Geçersiz seçenek! Tekrar deneyin.")
		}
	}
}

func listEvents() {
	if len(events) == 0 {
		fmt.Println("Henüz hiç etkinlik yok.")
		return
	}

	fmt.Println("Etkinlik Listesi:")
	for _, event := range events {
		fmt.Printf("ID: %d | Başlık: %s | Tarih: %s | Detaylar: %s\n", event.ID, event.Title, event.Date, event.Details)
	}
}

func addEvent() {
	lastEventID++
	fmt.Print("Etkinlik Başlığı: ")
	title := getUserInput()

	fmt.Print("Etkinlik Tarihi: ")
	date := getUserInput()

	fmt.Print("Etkinlik Detayları: ")
	details := getUserInput()

	events[lastEventID] = Event{ID: lastEventID, Title: title, Date: date, Details: details}
	fmt.Println("Etkinlik başarıyla eklendi!")
}

func editEvent() {
	fmt.Print("Düzenlemek istediğiniz etkinliğin ID'sini girin: ")
	eventID := getUserInputAsInt()

	event, found := events[eventID]
	if !found {
		fmt.Println("Belirtilen ID ile etkinlik bulunamadı.")
		return
	}

	fmt.Printf("Etkinlik Başlığı (%s): ", event.Title)
	title := getUserInputWithDefault(event.Title)

	fmt.Printf("Etkinlik Tarihi (%s): ", event.Date)
	date := getUserInputWithDefault(event.Date)

	fmt.Printf("Etkinlik Detayları (%s): ", event.Details)
	details := getUserInputWithDefault(event.Details)

	events[eventID] = Event{ID: eventID, Title: title, Date: date, Details: details}
	fmt.Println("Etkinlik başarıyla güncellendi!")
}

func deleteEvent() {
	fmt.Print("Silmek istediğiniz etkinliğin ID'sini girin: ")
	eventID := getUserInputAsInt()

	_, found := events[eventID]
	if !found {
		fmt.Println("Belirtilen ID ile etkinlik bulunamadı.")
		return
	}

	delete(events, eventID)
	fmt.Println("Etkinlik başarıyla silindi!")
}

func getUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getUserInputAsInt() int {
	input := getUserInput()
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Geçersiz girdi. Bir sayı bekleniyordu.")
		os.Exit(1)
	}
	return value
}

func getUserInputWithDefault(defaultValue string) string {
	input := getUserInput()
	if input == "" {
		return defaultValue
	}
	return input
}
