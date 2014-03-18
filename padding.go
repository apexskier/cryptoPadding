// Package cryptoPadding implements several standard byte padding schemes for
// use in block ciphers, such as standard AES.
//
// Byte padding appends bytes to data to make its total length a multiple of
// some block size.
//
// Example usage
//    import (
//        "github.com/apexskier/cryptoPadding"
//        "io/ioutil"
//    )
//    func main() {
//        var padding = cryptoPadding.AnsiX923
//        data, _ = ioutil.ReadFile("myfile")
//        paddedData = padding.Pad(data, 8)
//        unpaddedData = padding.Unpad(paddedData, 8)
//    }
package cryptoPadding

// BlockPadding represents an arbitrary byte padding scheme.
type BlockPadding interface {
	Pad(data []byte, blockSize int) (output []byte, err error)
	Unpad(data []byte, blockSize int) (output []byte, err error)
}

// padSize return the number of bytes needed to properly pad a size of data.
func padSize(dataSize, blockSize int) (ps int) {
	ps = blockSize - (dataSize % blockSize)
	if ps == 0 {
		ps = blockSize
	}
	return
}
