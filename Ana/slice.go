package main

import "fmt"

func main() {
	/*
		dizi := [5]int{1, 3, 5, 7, 9}
		var dilim []int = dizi[1:4]
		fmt.Println(dilim)
	*/
	a := [5]string{"Go", "Python", "Kotlin", "PHP", "JavaScript"}
	var slice []string = a[2:5]
	fmt.Println(slice)
}
