package logic2

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal2_5(n int) (result [][]int) {
	utils.CreateSlice(n)

	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 2
			} else {
				result[i][n-j-1] = start
				start += 2
			}
		}
	}
	return result
}
