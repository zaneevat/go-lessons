package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func main() {
	var snil []int
	fmt.Println(snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Println("slice is nil")
	}
	animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	a := animalsArr[0:2]
	fmt.Println(a)

	b := animalsArr[1:3]

	fmt.Println(b)

	animalSlice := []string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	fmt.Println(animalSlice)
}

func slices() {
	letters := []string{"a", "b", "c"}
	letters[0] = "1"
	letters = append(letters, "c")
	letters = append(letters, "d", "e", "f")

	fmt.Println(letters)

	emptyLetters := make([]string, 3)
	emptyLetters[0] = "1"
	emptyLetters[1] = "2"
	emptyLetters[2] = "3"
	emptyLetters = append(emptyLetters, "4")
	fmt.Println(emptyLetters)
}

func arrays() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Println(a[1])

	numbers := [...]int{3, 2, 3}
	numbers[2] = 4
}

func structs() {
	p2 := Point{
		X: 1,
		Y: 2,
		S: "string",
	}
	p2.method()
	fmt.Println(p2)
	fmt.Println(p2.X)
	p2.X = 123
	fmt.Println(p2)
}

func pointers() {
	a := "hello world"
	b := 42
	fmt.Println(a)
	fmt.Println(b)

	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = "ok"

	fmt.Println(a)
}
