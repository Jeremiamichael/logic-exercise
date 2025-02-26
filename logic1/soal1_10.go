package logic1

import "strconv"

func Soal1_10(n int) []string {
	var arr []string
	num := 2
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			arr = append(arr, strconv.Itoa(num))
			num *= 2
		} else {
			arr = append(arr, "fizz")
		}
	}
	return arr
}
