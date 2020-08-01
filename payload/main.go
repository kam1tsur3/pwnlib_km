package payload

import(
	"fmt"
)

func P32(n uint32) []byte{
	r := make([]byte, 4, 4)
	fmt.Println("P32(uint32)")
	return r
}

func U32(s []byte) uint32{
	n := 0
	fmt.Println("U32()")
	return n 
}

// for 32bit
func P64(n uint32) []byte{
	r := make([]byte, 8, 8)
	fmt.Println("P64(uint32)")
	return r
}

func P64(n uint64) []byte{
	r := make([]byte, 8, 8)
	fmt.Println("P64(uint64)")
	return r
}

func U64(s []byte) uint64{
	n := 0
	fmt.Println("U64()")
	return n
}
