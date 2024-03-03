package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	for i := 0; i <= len(arr)-1; i++ {
		fmt.Println(arr[i])
	}
}
