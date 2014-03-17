package crypto_padding

import (
    //"bytes"
    "crypto/rand"
    "errors"
    "fmt"
)

type ISO101226 struct {}

func (padding ISO101226) Pad(data []byte, blockSize int) (output []byte, err error) {
    if blockSize < 1 || blockSize >= 256 {
        return output, errors.New(fmt.Sprintf("blocksize is out of bounds: %v", blockSize))
    }
    var paddingBytes = padSize(len(data), blockSize)
    paddingSlice := make([]byte, paddingBytes - 1)
    _, err = rand.Read(paddingSlice)
    if err != nil {
        return output, err
    }
    paddingSlice = append(paddingSlice, byte(paddingBytes))
    fmt.Println(paddingSlice)
    output = append(data, paddingSlice...)
    return output, nil
}

func (padding ISO101226) Unpad(data []byte, blockSize int) (output []byte, err error) {
    var dataLen = len(data)
    if dataLen % blockSize != 0 {
        return output, errors.New("data's length isn't a multiple of blockSize")
    }
    var paddingBytes = int(data[dataLen - 1])
    if paddingBytes > blockSize || paddingBytes <= 0 {
        return output, errors.New(fmt.Sprintf("invalid padding found: %v", paddingBytes))
    }
    output = data[0:dataLen - paddingBytes]
    return output, nil
}
