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

	return r
}

func U32(s []byte) uint32{
	var n uint32 = 0

	if len(s) != 4{
		Local_Error("In U32(): the length of argument must be 4.\n")
	}

	for i, c := range s{
		n += uint32(c) << uint32(i*8)
	}

	return n
}

// for 64bit
func P64(n uint64) []byte{
	r := make([]byte, 8, 8)

	for i,_ := range r{
		r[i] = byte(n & 0xff)
		n >>= 8
	}

	return r
}

func U64(s []byte) uint64{
	var n uint64 = 0

	if len(s) != 8{
		Local_Error("In U64(): the length of argument must be 8.\n")
	}

	for i, c := range s{
		n += uint64(c) << uint32(i*8)
	}

	return n
}

// for house of corrosion
func Offset2Size(off int64) int64{
	return off*2-0x10
}

// Format String Bug
func Fsb(width, offset, loop, padding int64, data uint64) []byte{
	var r []byte
	var s string = ""
	var format string
	var mask uint32
	var write_n, write_all int64 = 0,padding
	move_up := 1 << (width*8)

	switch width {
		case 1:
			format = "$hhn"
			mask = 0xff
		case 2:
			format = "$hn"
			mask = 0xffff
		case 4:
			format = "$n"
			mask = 0xffffffff
		default:
			Local_Error("In Fsb(): width must be 1 or 2 or 4.\n")
	}
	for i := 0; i < loop; i++{
		write_n = ((data >> (i*8*width)) & mask) + move_up - (write_all & mask)
		if write_n > move_up { 
			write_n -= move_up
		}
		s += "%" + string(write_n) + "x%" + string(offset+i) + format
		write_all += write_n
	}
	r = []byte(s)
	return r
}
