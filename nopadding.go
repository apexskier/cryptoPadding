package cryptoPadding

// NoPadding fufils the padding interface to allow no padding.
type NoPadding struct{}

// Pad is a no-op
func (padding NoPadding) Pad(data []byte, blockSize int) (output []byte, err error) {
	return data, nil
}

// Unpad is a no-op
func (padding NoPadding) Unpad(data []byte, blockSize int) (output []byte, err error) {
	return data, nil
}
