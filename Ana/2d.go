package main

import "fmt"

func main() {
	/*
		var arr = [2][2]int{{0, 0}, {2, 4}}
		fmt.Println(arr)
		fmt.Println(arr[0][1])
		fmt.Println(arr[1][0])

	*/
	var arr = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 5}}
	fmt.Println(arr)
	fmt.Println(arr[3][0], arr[4][1])
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Arr[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}
}
