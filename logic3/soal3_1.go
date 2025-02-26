package logic3

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal3_1(n int) (result [][]int) {
	result = utils.CreateSlice(n)
	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i >= j {
				if i%2 == 0 {
					result[i][j] = num
				} else {
					result[i][i-j] = num
				}
				num += 2
			}
		}
	}
	return result
}
