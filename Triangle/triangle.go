package Triangle

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func triangle() {
	scanner:=bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome... First You Need To Choose The Type Of Triangle")
	fmt.Println("Isosceles Triangle? Equilateral Triangle? Scalene Triangle?")
	scanner.Scan()
	secim:=strings.ToLower(scanner.Text())
	
	fmt.Println("Area or Perimeter")
	scanner.Scan()
	secim1:=strings.ToLower(scanner.Text())
	
	fmt.Println("Please Enter First Side:")
	scanner.Scan()
	first,err:=strconv.ParseFloat(scanner.Text(),64)
	if err!=nil{
		log.Fatal(err)
	}
	
	fmt.Println("Please Enter Second Side:")
	scanner.Scan()
	second,err:=strconv.ParseFloat(scanner.Text(),64)
	if err!=nil{
		log.Fatal(err)
	}
	
	fmt.Println("Please Enter Third Side:")
	scanner.Scan()
	third,err:=strconv.ParseFloat(scanner.Text(),64)
	if err!=nil{
		log.Fatal(err)
	}
	
	if secim=="isosceles"{
		if secim1 == "area"{
			
		}
	}
	
}