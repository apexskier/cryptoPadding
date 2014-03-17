package crypto_padding

// Fufils the padding interface to allow no padding.
type NoPadding struct {}

func (padding NoPadding) Pad(data []byte, blockSize int) (output []byte, err error) {
    return data, nil
}

func (padding NoPadding) Unpad(data []byte, blockSize int) (output []byte, err error) {
    return data, nil
}
