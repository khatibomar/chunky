package main

import (
	"fmt"
)

func main() {
	var low int
	var mid int
	var high int

	var m map[int]int = map[int]int{
		0:  200,
		1:  200,
		2:  200,
		3:  200,
		4:  200,
		5:  200,
		6:  200,
		7:  200,
		8:  200,
		9:  200,
		10: 200,
		11: 200,
		12: 200,
		13: 404,
		14: 404,
		15: 404,
	}

	low = 0
	high = 40000
	currHigh := high

	for {
		high = high / 2
		if m[high] != 200 {
			currHigh = high
		} else {
			high = currHigh
			break
		}
	}
	fmt.Printf("starting with high = %d\n", high)
	for {
		mid = (high + low) / 2
		if m[mid] == 200 {
			low = mid
		} else {
			high = mid
		}
		if high == low+1 {
			break
		}
	}
	fmt.Println(low)
}
