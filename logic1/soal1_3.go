package logic1

func Soal1_3(n int) []int {
	result := make([]int, n)
	num := 3

	for i := 0; i < n; i++ {
		result[i] = num
		num += 3
	}
	return result
}
