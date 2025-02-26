package logic1

import (
	"strconv"
)

func Soal1_11(n int) []string {
	
	var arr []string
	num := 1
	first := true
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr = append(arr, "buzz")
		} else {
			arr = append(arr, strconv.Itoa(num))
			if first {
				num += 2
				first = false
			} else {
				num *= 2
			}
		}
	}
	return arr
}
