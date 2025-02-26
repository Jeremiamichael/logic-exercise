package logic2

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal2_4(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	start := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i%2 == 0 {
				result[i][j] = start
				start += 3
			} else {
				result[i][j] = start
				start += 3
			}
		}
	}
	return result
}
