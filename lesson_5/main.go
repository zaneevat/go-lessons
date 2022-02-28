package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) movePoint(x, y int) {
	p.X = x
	p.Y = y
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)
	p.movePoint(36, 44)
	fmt.Println(p)
	v := &p
	v.movePoint(1, 5)
	fmt.Println(v)
	fmt.Println(p)
}
