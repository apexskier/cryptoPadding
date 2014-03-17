package crypto_padding

import (
    "errors"
    "fmt"
    //"bytes"
)

type BlockPadding interface {
    Pad(data []byte, blocksize int) (output []byte, err error)
    Unpad(data []byte, blocksize int) (output []byte, err error)
}

// http://tools.ietf.org/html/rfc5652#section-6.3
type PKCS7 struct {}

func (padding PKCS7) Pad(data []byte, blocksize int) (output []byte, err error) {
    var paddingBytes = blocksize - (len(data) % blocksize)
    if paddingBytes == 0 {
        paddingBytes = blocksize
    }

    paddingSlice := make([]byte, paddingBytes)
    for i, _ := range paddingSlice {
        paddingSlice[i] = byte(paddingBytes)
    }

    output = append(data, paddingSlice...)

    return output, nil
}

func (padding PKCS7) Unpad(data []byte, blocksize int) (output []byte, err error) {
    var dataLen = len(data)
    if dataLen % blocksize != 0 {
        return output, errors.New("data's length isn't a multiple of blocksize")
    }
    var paddingBytes = int(data[dataLen - 1])
    if paddingBytes >= blocksize || paddingBytes <= 0 {
        return output, errors.New(fmt.Sprintf("invalid padding found: %v", paddingBytes))
    }
    var pad = data[dataLen - paddingBytes:dataLen - 1]
    for _, v := range pad {
        if int(v) != paddingBytes {
            return output, errors.New("invalid padding found")
        }
    }
    output = data[0:dataLen - paddingBytes]
    return output, nil
}
