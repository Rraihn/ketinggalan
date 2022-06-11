package main

import (
	array2 "generic/array"
)

func main() {
	Logic04Soal07(3)
}

func Logic04Soal07(n int) {
	m := (n * 2) + 1
	array := array2.NewNumberArray(n, m)
	a := 1
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if j%2 == 0 {
				array[i][j] = int32(a)
				a += 2
			} else if j%4 == 1 && i == 0 {
				array[i][j] = int32(a)
				a += 2
			} else if j%3 == 0 && i == n-1 {
				array[i][j] = int32(a)
				a += 2
			}
		}
	}
	// print array
	array2.PrintNumberArrayZeroEmpty(array)
}
