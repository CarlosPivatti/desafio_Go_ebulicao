package main

import "fmt"

func main() {
	x := soma(4, 2, 5)
	y := multiplicacao(10, 20)
	a := subtracao(10, 5)
	b := divisao(10, 2)
	fmt.Println(x, y, a, b)

}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func multiplicacao(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func subtracao(a int, i ...int) int {
	total := a
	for _, v := range i {
		total -= v
	}
	return total
}

func divisao(a int, i ...int) int {
	total := a
	for _, v := range i {
		total /= v
	}
	return total

}
