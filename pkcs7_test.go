package cryptoPadding

import (
    "bytes"
    "testing"
)

func TestPadUnpad_PKCS7(t *testing.T) {
    var p PKCS7
    testPadUnpad(t, p)
}

func TestPadCases_pkcs7(t *testing.T) {
    var p PKCS7
    data := []byte{21, 22, 23, 24, 25}

    padded, err := p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if !bytes.Equal(padded, []byte{21, 22, 23, 24, 25, 3, 3, 3}) {
        t.Fatalf("Wrong padding")
    }

    data = []byte{21, 22, 23, 24, 25, 26, 27, 28}

    padded, err = p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if !bytes.Equal(padded, []byte{21, 22, 23, 24, 25, 26, 27, 28, 8, 8, 8, 8, 8, 8, 8, 8}) {
        t.Fatalf("Wrong padding")
    }
}
