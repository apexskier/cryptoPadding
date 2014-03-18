package cryptoPadding

import (
    "testing"
)

func TestPadUnpad_iso10126(t *testing.T) {
    var p ISO10126
    testPadUnpad(t, p)
}

func TestPadCases_iso10126(t *testing.T) {
    var p ISO10126
    data := []byte{21, 22, 23, 24, 25}

    padded, err := p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if padded[7] != byte(3) {
        t.Fatalf("Wrong padding")
    }

    data = []byte{21, 22, 23, 24, 25, 26, 27, 28}

    padded, err = p.Pad(data, 8)
    if err != nil {
        t.Fatalf("Fail on Pad")
    }
    if padded[15] != byte(8) {
        t.Fatalf("Wrong padding")
    }
}
