package main

import "fmt"

type structureHere struct {
	N1, N2 int
}

func (s *structureHere) Sum() int {
	return s.N1 + s.N2
}

type interfaceHere interface {
	Sum() int
}

func main() {
	var a interfaceHere
	sh := structureHere{1, 2}
	os := otherStruct{2, 3}

	a = &sh
	fmt.Println(a.Sum())

	a = &os
	fmt.Println(a.Sum())

	var i interfaceHere = otherStruct{4, 5}
	fmt.Println(i.Sum())
}

type otherStruct struct {
	A, B int
}

func (o otherStruct) Sum() int {
	return o.A + o.B
}
