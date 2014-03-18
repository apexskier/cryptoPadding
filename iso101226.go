package cryptoPadding

import (
	"crypto/rand"
	"errors"
	"fmt"
)

// ISO10126 implements ISO 10126 byte padding. This has been withdrawn in 2007.
type ISO10126 struct{}

// Pad adds random bytes of padding. The last byte is the size of padding added.
//
// Example for a blocksize of 8:
//     -> [DD DD DD DD DD 23 138 03]
func (padding ISO10126) Pad(data []byte, blockSize int) (output []byte, err error) {
	if blockSize < 1 || blockSize >= 256 {
		return output, errors.New(fmt.Sprintf("blocksize is out of bounds: %v", blockSize))
	}
	var paddingBytes = padSize(len(data), blockSize)
	paddingSlice := make([]byte, paddingBytes-1)
	_, err = rand.Read(paddingSlice)
	if err != nil {
		return output, err
	}
	paddingSlice = append(paddingSlice, byte(paddingBytes))
	output = append(data, paddingSlice...)
	return output, nil
}


// Unpad removes padding.
func (padding ISO10126) Unpad(data []byte, blockSize int) (output []byte, err error) {
	var dataLen = len(data)
	if dataLen%blockSize != 0 {
		return output, errors.New("data's length isn't a multiple of blockSize")
	}
	var paddingBytes = int(data[dataLen-1])
	if paddingBytes > blockSize || paddingBytes <= 0 {
		return output, errors.New(fmt.Sprintf("invalid padding found: %v", paddingBytes))
	}
	output = data[0 : dataLen-paddingBytes]
	return output, nil
}
