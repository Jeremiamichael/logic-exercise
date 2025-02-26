package logic3

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal3_6(n int) (result [][]int) {
	result = utils.CreateSlice(n)
	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if i <= j {
				//result[i][j] = num
				if i%2 == 0 {
					result[i][j] = num
					result[n-1-i][j] = num
				} else {
					result[i][n-1-j] = num
					result[n-1-i][n-1-j] = num
				}
				num += 2
			}

		}
	}
	return result
}
