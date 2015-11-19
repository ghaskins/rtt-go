package main

import (
	"net"
	"os"
)

func NewUDS(payloadLen int) func() {

	laddr := "./rtt-go-plain.unix"

	// Start our listener in common-context so we don't race with the registration
	listener, err := net.Listen("unix", laddr)
	if err != nil {
		panic(err)
	}

	go func () {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		buf := make([]byte, payloadLen)

		for {
			conn.Read(buf)
			conn.Write(buf)
		}
	}()

	conn, err := net.Dial("unix", laddr)
	if err != nil {
		panic(err)
	}

	os.Remove(laddr)

	buf := make([]byte, payloadLen)

	return func() {
		conn.Write(buf)
		conn.Read(buf)
	}
}
