package slice

// use copy()
func Copy(b []byte) []byte {
	if b == nil {
		return nil
	}
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// append to a nil slice
func Append(b []byte) []byte {
	if b != nil && len(b) == 0 {
		return []byte{}
	}
	c := append([]byte(nil), b...)
	return c
}

// append to a pre-allocated slice
func Append2(b []byte) []byte {
	if b == nil {
		return nil
	}
	c := make([]byte, 0, len(b))
	c = append(c, b...)
	return c
}

func Append3(b []byte) []byte {
	c := append(b[:0:0], b...)
	return c
}
