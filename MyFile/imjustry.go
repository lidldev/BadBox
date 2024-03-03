package main

import "fmt"

const pi = 3.14

type daire struct {
	yaricap float64
}

func (d daire) alan() float64 {
	return pi * d.yaricap * d.yaricap
}

func main() {
	d1 := daire{5.0}
	fmt.Println(d1)

}
