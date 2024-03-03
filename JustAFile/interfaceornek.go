package main

import "fmt"

const pi = 3.14

type hesapla interface {
	alan() float64
}

type dare struct {
	yaricap float64
}

type dikdortgen struct {
	kisakenar float64
	uzunkenar float64
}

func (d dare) alan() float64 {
	return pi * d.yaricap * d.yaricap
}

func (g dikdortgen) alan() float64 {
	return g.uzunkenar * g.kisakenar
}

func hesapYap(h hesapla) {
	fmt.Println(h)
	fmt.Println(h.alan())
}

func main() {
	daire1 := dare{5.0}
	dikdortgen1 := dikdortgen{6.0, 9.0}

	hesapYap(daire1)
	hesapYap(dikdortgen1)
}
