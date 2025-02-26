package logic3

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal3_5(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		geser := 0
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				if i%2 == 1 {
					result[i][j] = num
				} else {
					result[i][n-1-geser] = num
					geser++
				}
			}
			num += 2
		}
	}
	return result
}
