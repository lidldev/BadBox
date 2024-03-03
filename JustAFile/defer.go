package main

import "fmt"

func start() {
	defer fmt.Println("PC AÇILDI")
	fmt.Println("PC AÇILIYOR...")
}
func close() {
	defer fmt.Println("PC KAPANDI")
	fmt.Println("PC KAPANIYOR...")
}

func main() {
	defer close()
	start()

}
