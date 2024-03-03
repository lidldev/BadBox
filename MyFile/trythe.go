package main

import (
	"fmt"
	"math/rand"
	"time"
)

type monster struct {
	name string
	heal uint
}

func main() {
	var opsion string

	m1 := monster{
		name: "Musho",
		heal: 100,
	}
	m2 := monster{
		name: "Sauha",
		heal: 100,
	}
	m3 := monster{
		name: "Naula",
		heal: 100,
	}
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(3) + 1

	if number == 1 {
		for {
			fmt.Println("Musho is attacking you what will you do")
			fmt.Scan(&opsion)
			if opsion == "attack" {
				downheal := rand.Intn(100) + 1
				fmt.Println(m1.heal - uint(downheal))
				m1.heal = m1.heal - uint(downheal)
				fmt.Println("Heal Left: ", m1.heal)
			}
			if opsion == "defence" {
				
			}
		}
	}
	if number == 2 {
		for {
			fmt.Println("Musho is attacking you what will you do")
			fmt.Scan(&opsion)
			if opsion == "attack" {
				downheal := rand.Intn(100) + 1
				fmt.Println(m2.heal - uint(downheal))
				m2.heal = m2.heal - uint(downheal)
				fmt.Println("Heal Left: ", m2.heal)
			}
		}
	}
	if number == 3 {
		for {
			fmt.Println("Musho is attacking you what will you do")
			fmt.Scan(&opsion)
			if opsion == "attack" {
				downheal := rand.Intn(100) + 1
				fmt.Println(m3.heal - uint(downheal))
				m3.heal = m3.heal - uint(downheal)
				fmt.Println("Heal Left: ", m3.heal)
			}
		}

	}
}
