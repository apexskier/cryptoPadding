package cryptoPadding

import (
    "testing"
    "bytes"
)

func TestPadUnpad_ansix923(t *testing.T) {
    var p AnsiX923
    testPadUnpad(t, p)
}

func TestPadCases_ansix923(t *testing.T) {
    var p AnsiX923
    data := []byte{21, 22, 23, 24, 25}

    padded, err := p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if !bytes.Equal(padded, []byte{21, 22, 23, 24, 25, 0, 0, 3}) {
        t.Fatalf("Wrong padding")
    }

    data = []byte{21, 22, 23, 24, 25, 26, 27, 28}

    padded, err = p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if !bytes.Equal(padded, []byte{21, 22, 23, 24, 25, 26, 27, 28, 0, 0, 0, 0, 0, 0, 0, 8}) {
        t.Fatalf("Wrong padding")
    }
}
