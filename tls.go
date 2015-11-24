package main

import (
	"crypto/tls"
	"net"
	"os"
)

type Connector interface {
	Listen(addr string) (net.Listener, error)
	Dial(addr string) (net.Conn, error)
}

//---------------------------------------
// TLS
//---------------------------------------
type tlsConnector struct{}

func (self *tlsConnector) Listen(addr string) (net.Listener, error) {
	cert, _ := tls.X509KeyPair([]byte(CertPEM), []byte(KeyPEM))

	config := &tls.Config{
		Certificates: make([]tls.Certificate, 1),
	}

	config.Certificates[0] = cert

	return tls.Listen("unix", addr, config)
}

func (self *tlsConnector) Dial(addr string) (net.Conn, error) {
	return tls.Dial("unix", addr, &tls.Config{InsecureSkipVerify: true})
}

//---------------------------------------
// UDS
//---------------------------------------
type udsConnector struct{}

func (self *udsConnector) Listen(addr string) (net.Listener, error) {
	return net.Listen("unix", addr)
}

func (self *udsConnector) Dial(addr string) (net.Conn, error) {
	return net.Dial("unix", addr)
}

//---------------------------------------
// Common Initializer
//---------------------------------------

func newConnector(connector Connector, addr string, payloadLen int) func() {

	// Start our listener in common-context so we don't race with the registration
	listener, err := connector.Listen(addr)
	if err != nil {
		panic(err)
	}

	go func() {
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

	conn, err := connector.Dial(addr)
	if err != nil {
		panic(err)
	}

	os.Remove(addr)

	buf := make([]byte, payloadLen)

	return func() {
		conn.Write(buf)
		conn.Read(buf)
	}
}

func NewTLS(payloadLen int) func() {
	return newConnector(&tlsConnector{}, "./rtt-go.tls", payloadLen)
}

func NewUDS(payloadLen int) func() {
	return newConnector(&udsConnector{}, "./rtt-go.uds", payloadLen)
}
