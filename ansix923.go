package cryptoPadding

import (
	"bytes"
	"errors"
	"fmt"
)

// AnsiX923 implements ANSI X.923 byte padding.
type AnsiX923 struct{}

// Pad adds padding of zeros, with the last byte being the padding size.
//
// Example for a blocksize of 8:
//     -> [DD DD DD DD DD 00 00 03]
//     -> [DD DD DD DD DD DD DD DD 00 00 00 00 00 00 00 08]
func (padding AnsiX923) Pad(data []byte, blockSize int) (output []byte, err error) {
	if blockSize < 1 || blockSize >= 256 {
		return output, errors.New(fmt.Sprintf("blocksize is out of bounds: %v", blockSize))
	}
	var paddingBytes = padSize(len(data), blockSize)
	paddingSlice := append(bytes.Repeat([]byte{byte(0)}, paddingBytes-1), byte(paddingBytes))
	output = append(data, paddingSlice...)
	return output, nil
}

// Unpad removes padding.
func (padding AnsiX923) Unpad(data []byte, blockSize int) (output []byte, err error) {
	var dataLen = len(data)
	if dataLen%blockSize != 0 {
		return output, errors.New("data's length isn't a multiple of blockSize")
	}
	var paddingBytes = int(data[dataLen-1])
	if paddingBytes > blockSize || paddingBytes <= 0 {
		return output, errors.New(fmt.Sprintf("invalid padding found: %v", paddingBytes))
	}
	var pad = data[dataLen-paddingBytes : dataLen-2]
	for _, v := range pad {
		if int(v) != 0 {
			return output, errors.New("invalid padding found")
		}
	}
	output = data[0 : dataLen-paddingBytes]
	return output, nil
}
