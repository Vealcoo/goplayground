package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var count int

type Test struct {
	queue chan int
	wait  sync.WaitGroup
}

func (t *Test) run() {
	for count < 10 {
		continue
	}

	for event := range t.queue {
		switch event {
		case 0:
			fmt.Println(0)
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		}
	}
	// for event := range t.queue {
	// 	switch event {
	// 	case 0:
	// 		fmt.Println(0)
	// 	case 1:
	// 		fmt.Println(1)
	// 	case 2:
	// 		fmt.Println(2)
	// 	}
	// }
}

func (t *Test) serve() {
	tk := time.Tick(1 * time.Second)
	for range tk {
		fmt.Println("<-")
		r := rand.Intn(3)
		t.queue <- r

		count++

		if count == 5 {
			fmt.Println("length:", len(t.queue))
			fmt.Println(fmt.Sprintf("%x", t.queue))
		}
	}
}

func Run() {
	t := new(Test)
	t.queue = make(chan int, 100)
	go t.run()
	t.serve()
}
