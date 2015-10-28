package main

import (
	"fmt"
	"flag"
	"time"
	"crypto/aes"
)

type Test struct {
	Name string
	Func func()
}

var aesKey = []byte("12345678901234567890123456789012")

func NewAES() func() {
	cipher, _ := aes.NewCipher(aesKey)
	output := make([]byte, len(aesKey))

	return func() { cipher.Encrypt(output, aesKey) /* just re-use aesKey as plain-text, who cares */}
}

func main() {
	iterations := *flag.Int("iterations", 10000, "the number of iterations per test")

	flag.Parse();

	tests := []Test{
		Test{"null", func() {} },
		//Test{"timer validation", func() { time.Sleep(100 * time.Microsecond)} },
		Test{"AES", NewAES() },
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
