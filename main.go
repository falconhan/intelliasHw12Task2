package main

import "fmt"

func main() {

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	ch := make(chan int)
	go func(ch chan int) { ch <- Sum(n[0]) }(ch)
	go func(ch chan int) { ch <- Sum(n[1]) }(ch)
	go func(ch chan int) { ch <- Sum(n[2]) }(ch)

	v1 := <-ch
	v2 := <-ch
	v3 := <-ch

	res := v1 + v2 + v3

	fmt.Println(res)

}

func Sum(n []int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum = sum + n[i]
	}
	return sum
}
