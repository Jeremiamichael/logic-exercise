package logic1

func Soal1_7(n int) []int {
	result := make([]int, n)

	mid := n / 2
	num := 1

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

func Soal1_7b(n int) []int {
	result := make([]int, n)

	mid := n / 2
	num := 1

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
