package _6

// my version
func PlusOne(digits []int) []int {
	len := len(digits)
	carry := false
	if digits[len-1] == 9 {
		carry = true
		digits[len-1] = 0
	} else {
		digits[len-1]++
		return digits
	}

	for i := len - 2; i > -1; i-- {
		if carry == false {
			break
		}
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			carry = false
		}
	}

	if carry != false {
		digits[0] = 0
		return append([]int{1}, digits...)
	}
	return digits
}

// reference
func PlusOne2(digits []int) []int {
	var n int = len(digits)
	for i := n - 1; i > -1; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}
	var a = make([]int, n+1)
	a[0] = 1
	return a

}
