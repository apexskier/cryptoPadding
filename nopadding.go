package cryptoPadding

// Fufils the padding interface to allow no padding.
type NoPadding struct{}

// No-op
func (padding NoPadding) Pad(data []byte, blockSize int) (output []byte, err error) {
	return data, nil
}

// No-op
func (padding NoPadding) Unpad(data []byte, blockSize int) (output []byte, err error) {
	return data, nil
}
