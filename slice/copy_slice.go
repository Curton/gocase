package slice

// use copy()
func Copy(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// append to a nil slice
func Append(b []byte) []byte {
	var c []byte
	c = append(c, b...)
	return c
}

// append to a pre-allocated slice
func Append2(b []byte) []byte {
	c := make([]byte, 0, len(b))
	c = append(c, b...)
	return c
}
