package logic3

import "github.com/Jeremiamichael/logic-exercise/utils"

func Soal3_3(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	num := 1
	for i := 0; i < n; i++ {
		for j := num; j < n; j++ {
			if i <= j {
				if i == 0 {
					result[i][j-i] = num
					num += 3
				} else if i%2 == 0 {
					result[i][j-i] = num
					num += 2
					if j == n-1 {
						num += 1
					}
				} else {
					result[i][n-j-1] = num
					num += 3
					if j == n-1 {
						num -= 1
					}
				}
			}
		}
	}
	return result
}
