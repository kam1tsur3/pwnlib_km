package network

import (
	"fmt"
	"os"
)

func Local_Error(s string) {
	fmt.Fprintf(os.Stderr, s)
	os.Exit(1)
}

func Test() {
	fmt.Println("in connection.go")
}
// Connection
/*
type Connection struct {

}

func (conn Connection) Send(bs []byte){

}

func (conn Connection)

*/
