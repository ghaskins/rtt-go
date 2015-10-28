package main

import (
	"fmt"
	"flag"
	"time"
	"crypto/aes"
	"crypto/rand"
)

type Test struct {
	Name string
	Func func()
}

func NewRand(len int) []byte {
	data := make([]byte, len)
	actual, err := rand.Read(data)

	if err != nil {
		panic(err.Error())
	}

	if actual != len {
		panic(fmt.Sprintf("short read"))
	}

	return data
}

func NewAES(payloadLen int) func() {
	cipher, _ := aes.NewCipher(NewRand(32))
	input := NewRand(payloadLen)
	output := make([]byte, payloadLen)

	return func() { cipher.Encrypt(output, input) }
}

func main() {
	iterations := *flag.Int("iterations", 10000, "the number of iterations per test")
	payloadLen := *flag.Int("payload", 256 * 1024, "the size of the payload to use")

	flag.Parse();

	tests := []Test{
		Test{"null", func() {} },
		//Test{"timer validation", func() { time.Sleep(100 * time.Microsecond)} },
		Test{"AES", NewAES(payloadLen) },
	}

	fmt.Printf("iterations: %d\n", iterations)

	for _, test := range tests {

		fmt.Print("Running test \"" + test.Name + "\"...")

		t0 := time.Now()
		for i := 0; i<iterations; i++ {
			test.Func()
		}
		t1 := time.Now()

		fmt.Printf("done: %dns/iteration\n", int(t1.Sub(t0).Nanoseconds())/iterations)
	}

}
