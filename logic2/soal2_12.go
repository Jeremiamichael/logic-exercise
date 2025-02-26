package logic2

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal2_12(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	for row := 0; row < n; row++ {
		num := 1
		for col := 0; col < n; col++ {
			if row >= col && row+col <= n-1 {
				result[row][col] = num
			} else if row <= col && row+col >= n-1 {
				result[row][col] = num
			}
			num += 2
		}
	}
	return result
}
