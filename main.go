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
	s,t := swap("hello","world")
	fmt.Println(s,t)
	p,q := split(100)
	fmt.Println("1st Number : ",p)
	fmt.Println("2nd number : ",q)
}

func GenrateRandInt(r int) int {
	randInt := rand.Intn(r)

	return randInt
}
func GenerateRandFloat() float64 {
	randFloat := rand.Float64()

	return randFloat
}

func swap (x,y string) (string , string) {
	return y,x
}
func split(sum int) (x, y int) {
	x = sum *4 /9
	y = sum - x
	return
}