package main

import "fmt"

const d string = "name"

func main() {
	defer fmt.Println("world")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	for sum < 200 {
		sum += 2
	}

	for sum < 1000 {
		sum += 10
	}

	fmt.Println(sum)

	a := 0
	for a < 1000 {
		i := isTest(a)
		if i == 1 {
			fmt.Println("a = 2")
		} else if i == 3 {
			fmt.Println("a = 3")
		} else {
			fmt.Printf("unknown a. a=%d", i)
		}

		switch i {
		case 1:
			fmt.Println("a = 2")
		case 2:
			fmt.Println("a = 3")
		default:
			fmt.Printf("unknown a, a=%d", i)
		}
		a++
	}
	fmt.Println("hello ")
}

func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 2
	}

	return 3
}
