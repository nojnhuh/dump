package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first_num := 0
	second_num := 1
	x := 0
	return func(x, first_num, second_num int) int
	{
		if x == 0 {
			return first_num
		}
		if x == 1 {
			return second_num
		}
		if x % 2 == 0 {
			first_num = first_num + second_num
			return first_num
		}
		second_num = second_num + first_num
		return second_num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}