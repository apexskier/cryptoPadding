package cryptoPadding

import (
	"bytes"
	"errors"
	"fmt"
)

// ZeroPadding adds padding of zeros.
type ZeroPadding struct{}

// Pad adds paddings of all zeros.
//
// Example for a blocksize of 8:
//     -> [DD DD DD DD DD 00 00 00]
func (padding ZeroPadding) Pad(data []byte, blockSize int) (output []byte, err error) {
	if blockSize < 1 || blockSize >= 256 {
		return output, fmt.Errorf("blocksize is out of bounds: %v", blockSize)
	}
	var paddingBytes = padSize(len(data), blockSize)
	paddingSlice := bytes.Repeat([]byte{byte(0)}, paddingBytes)
	output = append(data, paddingSlice...)
	return output, nil
}

// Unpad attempts to remove padding.
//
// Will not behave properly if the last character of the unpadded data is a zero.
func (padding ZeroPadding) Unpad(data []byte, blockSize int) (output []byte, err error) {
	var dataLen = len(data)
	if dataLen%blockSize != 0 {
		return output, errors.New("data's length isn't a multiple of blockSize")
	}
	var paddingBytes = 0
	for data[dataLen-1-paddingBytes] == 0 {
		paddingBytes++
	}
	if paddingBytes > blockSize || paddingBytes <= 0 {
		return output, fmt.Errorf("invalid padding found: %v", paddingBytes)
	}
	output = data[0 : dataLen-paddingBytes]
	return output, nil
}
