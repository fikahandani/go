package main

import "fmt"

func main() {
	fmt.Println("Closure adalah sebuah fungsi yang disimpan dalam sebuah variable")
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < e:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 1, 5, 6, 8, 7, 9, 4, 5}
	var min, max = getMinMax(numbers)
	fmt.Printf("Data: %v\nMin : %v\nMax : %v\n", numbers, min, max)
}
