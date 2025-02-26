package logic1

func Soal1_8(n int) []int {
	mid := n / 2
	num := 2
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 2
		} else if i > mid {
			num -= 2
		}
	}
	return result
}

func Soal1_8b(n int) []int {
	mid := n / 2
	num := 2
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = num
		if i < mid {
			num += 2
		} else if i >= mid {
			num -= 2
		}
	}
	return result
}
