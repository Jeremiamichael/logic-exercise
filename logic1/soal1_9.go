package logic1

func Soal1_9(n int) []int {
	result := make([]int, n)
	mid := n / 2
	num := 2

	for i := 0; i < mid; i++ {
		result[i] = num
		if i < mid {
			num += 3
		} else if i > mid {
			num -= 3
		}
	}
	return result
}

func Soal1_9b(n int) []int {
	result := make([]int, n)
	mid := n / 2
	num := 2
	for i := 0; i < mid; i++ {
		result[i] = num
		if i < mid {
			num += 3
		} else if i >= mid {
			num -= 3
		}
	}
	return result
}
