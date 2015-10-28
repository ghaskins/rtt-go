package main

import (
	"fmt"
	"flag"
	"time"
)

type Test struct {
	Name string
	Func func()
}

func main() {
	iterations := *flag.Int("iterations", 10000, "the number of iterations per test")

	flag.Parse();

	tests := []Test{
		Test{"null", func() {} },
		//Test{"timer validation", func() { time.Sleep(100 * time.Microsecond)} },
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
