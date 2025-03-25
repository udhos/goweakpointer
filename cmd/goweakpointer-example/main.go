// Package main implements the tool.
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"time"
	"weak"
)

type info struct {
	value atomic.Int64
	data  [100_000_000]byte
}

func main() {

	var wg sync.WaitGroup

	{
		value := new(info)

		weakValue := weak.Make(value)

		wg.Add(1)

		go worker(&wg, weakValue)

		const count = 5

		for i := range count {
			fmt.Printf("main   : value: %2d (count down: %2d)\n",
				value.value.Add(1), count-i)
			sleep()
		}

		fmt.Printf("main   : finished\n")
	}

	fmt.Printf("main   : waiting worker\n")
	wg.Wait()
	fmt.Printf("main   : waiting worker - done\n")
}

func worker(wg *sync.WaitGroup, weakValue weak.Pointer[info]) {
	fmt.Printf("worker : want to work forever\n")
	for {
		strong := weakValue.Value()
		if strong == nil {
			fmt.Printf("worker : lost the pointer!\n")
			break
		}
		fmt.Printf("worker : value: %2d\n", strong.value.Add(1))
		strong = nil

		//
		// put pressure on the GC
		//
		x := new(info)
		x.data[0] = 'a'
		x = nil

		sleep()
	}
	fmt.Printf("worker : won't work forever anymore, exiting to release resources\n")
	wg.Done()
}

func sleep() {
	time.Sleep(time.Duration(rand.IntN(1000))*time.Millisecond + time.Second)
}
