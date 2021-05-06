package main

import "fmt"

func main() {
	fmt.Println("result -> ", fibonacci(20))
}

func fibonacci(n int) int {
	fmt.Println("Input: ", n)

	x, y := 0, 1

	if n == 1 {
		return x
	}

	fmt.Println("1 -> ", x)

	if n == 2 {
		return y
	}

	fmt.Println("2 -> ", y)

	for i := 3; i <= n; i++ {
		x, y = y, x+y
		fmt.Println(i, " -> ", y)
	}

	return y
}
