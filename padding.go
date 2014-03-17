package crypto_padding

import (
)

type BlockPadding interface {
    Pad(data []byte, blockSize int) (output []byte, err error)
    Unpad(data []byte, blockSize int) (output []byte, err error)
}

func padSize(dataSize, blockSize int) (ps int) {
    ps = blockSize - (dataSize % blockSize)
    if ps == 0 {ps = blockSize}
    return
}
