package logic2

func Soal2_3(n int) (result [][]int) {
	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 2
			} else {
				result[i][j] = start
				start += 2
			}
		}
	}
	return result
}
