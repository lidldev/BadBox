package main

import "fmt"

func main() {
	/*
		dizi := [5]int{1, 3, 5, 7, 9}
		var dilim []int = dizi[1:4]
		fmt.Println(dilim)
	*/
	dizi := [5]string{"Go", "Python", "Kotlin", "PHP", "JavaScript"}
	var dilim []string = dizi[1:4]
	fmt.Println("Slice:", dilim)
	fmt.Println("Array:", dizi)
}
