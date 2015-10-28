package main

import (
	"fmt"
	"flag"
	"time"
	"crypto/aes"
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
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

func NewSHA256(payloadLen int) func() {
	input := NewRand(payloadLen)

	return func() { sha256.Sum256(input) }
}

func NewECDSA(payloadLen int) func() {
	pubkeyCurve := elliptic.P384()

	privatekey := new(ecdsa.PrivateKey)
	privatekey, _ = ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
	input := NewRand(payloadLen)
	digestA := sha256.Sum256(input)
	sigA, sigB, _ := ecdsa.Sign(rand.Reader, privatekey, digestA[:])

	return func() {
			digestB := sha256.Sum256(input)
			verified := ecdsa.Verify(&privatekey.PublicKey, digestB[:], sigA, sigB)
			if verified == false {
				panic("signature verification failed")
			}
		}
}

func main() {
	iterations := flag.Int("iterations", 100, "the number of iterations per test")
	payloadLen := flag.Int("payload", 1 * 1024, "the size of the payload to use")

	flag.Parse();

	tests := []Test{
		Test{"null", func() {} },
		//Test{"timer validation", func() { time.Sleep(100 * time.Microsecond)} },
		Test{"AES", NewAES(*payloadLen) },
		Test{"SHA256", NewSHA256(*payloadLen) },
		Test{"ECDSA verify", NewECDSA(*payloadLen) },
	}

	fmt.Printf("iterations: %d payloadLen: %d\n", *iterations, *payloadLen)

	for _, test := range tests {

		fmt.Print("Running test \"" + test.Name + "\"...")

		t0 := time.Now()
		for i := 0; i<*iterations; i++ {
			test.Func()
		}
		t1 := time.Now()

		fmt.Printf("done: %dns/iteration\n", int(t1.Sub(t0).Nanoseconds())/(*iterations))
	}

}
