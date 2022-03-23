package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jari(), 2)
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.jari()
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

func main() {
	var b hitung

	b = lingkaran{14.0}

	fmt.Println(b.luas())
	fmt.Println(b.keliling())
	fmt.Println(b.(lingkaran).jari())

	b = persegi{10.0}
	fmt.Println(b.luas())
	fmt.Println(b.keliling())
}
