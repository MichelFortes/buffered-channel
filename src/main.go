package main

import (
	"context"
	"fmt"
	"michelfortes/batch/worker"
)

func main() {

	ch := make(chan int, 2)
	ctx := context.Background()

	go worker.StartWorker(ctx, "worker 1", ch)
	go worker.StartWorker(ctx, "worker 2", ch)

	for i := 0; i < 100; i++ {
		fmt.Printf("sending %d\n", i)
		ch <- i
	}

}
