package main

import (
	"fmt"
)

func calculateHammingWeight(x uint32) uint32 {
	x = ((x & 0xAAAAAAAA) >> 1) + (x & 0x55555555)
	x = ((x & 0xCCCCCCCC) >> 2) + (x & 0x33333333)
	x = ((x & 0xF0F0F0F0) >> 4) + (x & 0x0F0F0F0F)
	x = ((x & 0xFF00FF00) >> 8) + (x & 0x00FF00FF)
	x = ((x & 0xFFFF0000) >> 16) + (x & 0x0000FFFF)
	return x
}

func next(x uint32) uint32 {
	var y uint32
	var c uint32

	y = x & -x
	c = x + y
	x = (((x ^ c) >> 2) / y) | c

	return x
}

func main() {
	var curr uint32

	fmt.Print("Enter first element: ")
	fmt.Scanf("%d", &curr)

	upper_bound := uint32(0x1000000)

	for ; curr < upper_bound; curr = next(curr) {
		fmt.Printf("value = %#x (%d), hamming weight = %d\n", curr, curr, calculateHammingWeight(curr))
	}
}
