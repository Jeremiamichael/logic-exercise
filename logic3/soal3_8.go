package logic3

func SOal3_8(n int) (result [][]int) {

	mid := (n - 1) / 2
	num := 1
	for k := 0; k <= mid; k++ {
		for b := 0; b < n; b++ {
			if k+b >= mid && b-k <= mid {
				if k%2 == 1 {
					result[b][k] = num
					result[b][n-1-k] = num
				} else {
					result[n-1-b][k] = num
					result[n-1-b][n-1-k] = num
				}
				num += 2
			}
		}
	}
	return result
}
