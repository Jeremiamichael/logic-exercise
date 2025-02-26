package logic1

func Soal1_12(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (2*(i%4) + 1))
	}
	return arr
}
