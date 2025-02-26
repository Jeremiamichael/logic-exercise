package logic1

func Soal1_5(n int) []int {
	result := make([]int, n)
	num := 20
	for i := 0; i < n; i++ {
		result[i] = num
		num -= 2
	}
	return result
}
