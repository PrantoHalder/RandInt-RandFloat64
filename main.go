package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//addition random int
	a := GenrateRandInt(10)
	fmt.Println(a)
	b := GenrateRandInt(10)
	fmt.Println(b)
	c := a + b
	fmt.Println(c)

	//adition random float
	d := GenerateRandFloat()
	fmt.Println(d)
	e := GenerateRandFloat()
	fmt.Println(e)
	f := d + e
	fmt.Println(f)
}

func GenrateRandInt(r int) int {
	randInt := rand.Intn(r)

	return randInt
}
func GenerateRandFloat() float64 {
	randFloat := rand.Float64()

	return randFloat
}
