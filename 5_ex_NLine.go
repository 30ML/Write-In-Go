package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		pricePant := i * 10
		priceSword := i*10 + 10
		fmt.Println("타잔이 " + pricePant + "원짜리 팬티를 입고, " + priceSword + "원짜리 칼을 차고 노래를 한다.")
	}
}
