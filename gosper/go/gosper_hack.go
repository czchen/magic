package main

import (
	"fmt"
)

func calculateHammingWeight(x uint64) uint64 {
	x = ((x & 0xAAAAAAAAAAAAAAAA) >> 1) + (x & 0x5555555555555555)
	x = ((x & 0xCCCCCCCCCCCCCCCC) >> 2) + (x & 0x3333333333333333)
	x = ((x & 0xF0F0F0F0F0F0F0F0) >> 4) + (x & 0x0F0F0F0F0F0F0F0F)
	x = ((x & 0xFF00FF00FF00FF00) >> 8) + (x & 0x00FF00FF00FF00FF)
	x = ((x & 0xFFFF0000FFFF0000) >> 16) + (x & 0x0000FFFF0000FFFF)
	x = ((x & 0xFFFFFFFF00000000) >> 32) + (x & 0x00000000FFFFFFFF)
	return x
}

func next(x uint64) uint64 {
	var y uint64
	var c uint64

	y = x & -x
	c = x + y
	x = (((x ^ c) >> 2) / y) | c

	return x
}

func main() {
	var curr uint64

	fmt.Print("Enter first element: ")
	fmt.Scanf("%d", &curr)

	upper_bound := uint64(0x100000000000000)

	for ; curr < upper_bound; curr = next(curr) {
		fmt.Printf("value = %#x (%d), hamming weight = %d\n", curr, curr, calculateHammingWeight(curr))
	}
}
