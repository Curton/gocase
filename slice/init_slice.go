package slice

func InitSlice() {
	a := make([]int, 50, 100)
	b := new([100]int)[0:50]
	_, _ = a, b
}
