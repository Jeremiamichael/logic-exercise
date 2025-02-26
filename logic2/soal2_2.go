package logic2

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal2_2(n int) (result [][]int) {
	result = utils.CreateSlice(n)
	for i := 0; i < n; i++ {
		num := 2
		for j := 0; j < n; j++ {
			result[i][j] = num
			num += 2
		}
	}
	return result
}
