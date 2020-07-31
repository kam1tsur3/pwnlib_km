package payload

import(
	"fmt"
)

func P32(n uint32) []byte{
	fmt.Println("P32(uint32)")
}

func U32(s []byte) uint32{
	fmt.Println("U32()")
}

// for 32bit
func P64(n uint32) []byte{
	fmt.Println("P64(uint32)")
}

func P64(n uint64) []byte{
	fmt.Println("P64(uint64)")
}

func U64(s []byte) uint64{
	fmt.Println("U64()")
}
