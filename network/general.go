package network

import (
	"fmt"
	"os"
)

func Local_Error(s string) {
	fmt.Fprintf(os.Stderr, s)
	fmt.Println("")
	os.Exit(1)
}
