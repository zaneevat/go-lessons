package main

import "fmt"

func totalPrice(intPrice int) func(int) int {
	sum := intPrice
	return func(a int) int {
		sum += a
		return sum
	}
}

func main() {
	orderPrice := totalPrice(1)
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
}

func doSomething(callback func(int, int) int, s string) int {
	result := callback(5, 1)
	fmt.Println(s)
	return result
}

func callback() {
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}

	result := doSomething(sumCallback, "plus")
	fmt.Println(result)

	multipleCallback := func(n1, n2 int) int {
		return n1 * n2
	}

	result = doSomething(multipleCallback, "multiple")
	fmt.Println(result)
}
