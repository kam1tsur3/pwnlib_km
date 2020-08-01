package payload

import(
	"fmt"
	"os"
)

func Local_Error(s string) {
	fmt.Fprintf(os.Stderr, s)
	os.Exit(1)
}

// for 32bit
func P32(n uint32) []byte{
	r := make([]byte, 4, 4)

	for i,_ := range r{
		r[i] = byte(n & 0xff)
		n >>= 8
	}

	fmt.Println("P32(uint32)")
	return r
}
/*
func U32(s []byte) uint32{
	var n uint32 = 0

	if len(s) != 4{
		Local_Error("In U32(): the length of argument must be 4.\n")
	}

	for i, c := range s{
		n += uint32(c) << i*8
	}

	fmt.Println("U32()")
	return n
}
*/

// for 64bit
/*func P64(n uint32) []byte{
	r := make([]byte, 8, 8)
	fmt.Println("P64(uint32)")
	return r
}
*/
/*
func P64(n uint64) []byte{
	r := make([]byte, 8, 8)

	for i,_ := range r{
		r[i] = n && uint64(0xff)
		n >>= 8
	}

	fmt.Println("P64(uint64)")
	return r
}

func U64(s []byte) uint64{
	var n uint64 = 0

	if len(s) != 8{
		Local_Error("In U64(): the length of argument must be 8.\n")
	}

	for i, c := range s{
		n += uint64(c) << (i*8)
	}

	fmt.Println("U64()")
	return n
}
*/
