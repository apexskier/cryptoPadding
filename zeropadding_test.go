package cryptoPadding

import (
    "bytes"
    "testing"
)

func TestPadCases_zero(t *testing.T) {
    var p ZeroPadding
    data := []byte{21, 22, 23, 24, 25}

    padded, err := p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if !bytes.Equal(padded, []byte{21, 22, 23, 24, 25, 0, 0, 0}) {
        t.Fatalf("Wrong padding")
    }

    data = []byte{21, 22, 23, 24, 25, 26, 27, 28}

    padded, err = p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if !bytes.Equal(padded, []byte{21, 22, 23, 24, 25, 26, 27, 28, 0, 0, 0, 0, 0, 0, 0, 0}) {
        t.Fatalf("Wrong padding")
    }
}
