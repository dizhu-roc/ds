package main

import (
	"fmt"

	"github.com/dizhu-roc/ds/loopqueue"
)

func main() {
	var q = loopqueue.New(200)
	fmt.Println(q)

	for i := 1; i <= 100; i++ {
		q.EnQueue(i)
	}
	fmt.Println(q)

	for i := 0; i < 10; i++ {
		q.DeQueue()
	}
	fmt.Println(q)

	fmt.Println(q.GetFront())
}
