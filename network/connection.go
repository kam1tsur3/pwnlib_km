package network

import (
	"fmt"
	"net"
)

// Connection
type PwnConn struct {
	conn *net.Conn
}

func NewConn(address string) *PwnConn{
	var pwnconn PwnConn
	conn, err := net.Dial("tcp", address)
	if err != nil {
		//Local_Error("In NewConn: net.Dial Error\n")
		fmt.Println("In NewConn: net.Dial Error\n")
	}
	pwnconn.conn = &conn
	return &pwnconn 
}

func (pc *PwnConn) Recv() []byte{
	
}

/*
func (pc *PwnConn) Send(bs []byte){

}

/*

func (pc *PwnConn) Recvline() []byte{

}

func (pc *PwnConn) Recvuntil() []byte{

}

func (pc *PwnConn) Sendline(bs []byte){

}

func (pc *PwnConn) Sendafter(s string, bs []byte){

}

func (pc *PwnConn) Sendlineafter(s string, bs []byte){

}

*/
