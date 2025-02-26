package logic2

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal2_10(n int) (result [][]int) {
	result = utils.CreateSlice(n)
	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				result[i][j] = num
			} else if i+j == n-1 {
				result[j][i] = num
			}
		}
		num += 2
	}
	return result
}
