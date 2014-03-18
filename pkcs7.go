package cryptoPadding

import (
	"bytes"
	"errors"
	"fmt"
)

// PKCS7 implements PKCS#7 padding, described at
// http://tools.ietf.org/html/rfc5652#section-6.3.
type PKCS7 struct{}

// Pad adds padding, with each padded byte being the total number of bytes
// added.
//
// Example for a blocksize of 8:
//     -> [DD DD DD DD DD 03 03 03]
func (padding PKCS7) Pad(data []byte, blockSize int) (output []byte, err error) {
	if blockSize < 1 || blockSize >= 256 {
		return output, fmt.Errorf("blocksize is out of bounds: %v", blockSize)
	}
	var paddingBytes = padSize(len(data), blockSize)
	paddingSlice := bytes.Repeat([]byte{byte(paddingBytes)}, paddingBytes)
	output = append(data, paddingSlice...)
	return output, nil
}

// Unpad removes padding.
func (padding PKCS7) Unpad(data []byte, blockSize int) (output []byte, err error) {
	var dataLen = len(data)
	if dataLen%blockSize != 0 {
		return output, errors.New("data's length isn't a multiple of blockSize")
	}
	var paddingBytes = int(data[dataLen-1])
	if paddingBytes > blockSize || paddingBytes <= 0 {
		return output, fmt.Errorf("invalid padding found: %v", paddingBytes)
	}
	var pad = data[dataLen-paddingBytes : dataLen-1]
	for _, v := range pad {
		if int(v) != paddingBytes {
			return output, errors.New("invalid padding found")
		}
	}
	output = data[0 : dataLen-paddingBytes]
	return output, nil
}
