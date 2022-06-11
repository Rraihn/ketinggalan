package main

import array2 "generic/array"

func main() {
	Logic04Soal05(4)
}

func Logic04Soal05(n int) {
	array := array2.NewNumberArray(n, n)
	a := 1

	for i := 1; i <= n; i++ {
		for j := 1; j < n; j++ {
			if i == 1 {
				array[n-1-j][0] = int32(a)
			} else if i == 2 {
				array[0][j] = int32(a)
			} else if i == 3 {
				array[j][n-1] = int32(a)
			} else if i == 4 {
				array[n-1][n-1-j] = int32(a)
			}
			a++
		}
	}
	array2.PrintNumberArrayZeroEmpty(array)
}
