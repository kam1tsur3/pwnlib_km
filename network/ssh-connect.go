package network

import (
	"fmt"
	"os"
)

func Local_Error(s string) {
	fmt.Fprintf(os.Stderr, s)
	os.Exit(1)
}

func Test2() {
	fmt.Println("in ssh-connect.go")
}
/*
type SSH_Connection struct {

}

func
*/