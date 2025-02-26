package logic2

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal2_11(n int) (result [][]int) {
	utils.CreateSlice(n)

	for i := 0; i < n; i++ {
		num := 1
		for j := 0; j < n; j++ {
			if i <= j {
				result[i][j] = num
			}
			num += 2
		}
	}
	return result
}
