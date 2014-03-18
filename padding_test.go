package cryptoPadding

import (
    "testing"
    "crypto/rand"
    "bytes"
)

func TestPadsize(t *testing.T) {
    for blocksize := 0; blocksize < 32; blocksize++ {
        for i := 1; i < blocksize; i++ {
            if ps := padSize(i, blocksize); (ps+i) % blocksize != 0 {
                t.Fatalf("padsize shouldn't have returned %v for dataSize of %v and blockSize of %v", ps, i, blocksize)
            }
        }
    }
}

func testPadUnpad(t *testing.T, p BlockPadding) {
    for blocksize := 1; blocksize < 128; blocksize++ {
        for i := 1; i < blocksize * 2; i++ {
            data := make([]byte, i)
            _, err := rand.Read(data)
            if err != nil { panic(err) }

            padded, err := p.Pad(data, blocksize)
            if err != nil {
                t.Fatalf("Fail on Pad for %v", i)
            }

            unpadded, err := p.Unpad(padded, blocksize)
            if err != nil {
                t.Fatalf("fail on Unpad for %v", i)
            }
            if !bytes.Equal(unpadded, data) {
                t.Fatalf("unpadded not equal to original for %v", i)
            }
        }
    }
}
