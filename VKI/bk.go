package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Kilonuzu Yazın:")
	scanner.Scan()
	kilo, _ := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("Boyunuzu Yazın:")
	scanner.Scan()
	boy, _ := strconv.ParseFloat(scanner.Text(), 64)
	vki := kilo / (boy * boy)
	fmt.Printf("Vücüt Kitle İndeksiniz: %f", vki)

}
