package worker

import (
	"context"
	"fmt"
	"time"
)

var workers map[string]*worker

type worker struct {
	ctx  context.Context
	name string
	ch   chan int
}

func (w *worker) run() {
	fmt.Printf("[ %s ] worker started\n", w.name)
	for {
		select {
		case x := <-w.ch:
			fmt.Printf("[ %s ] processing id %d\n", w.name, x)
			time.Sleep(2 * time.Second)
		case <-w.ctx.Done():
			fmt.Printf("[ %s ] context finished\n", w.name)
		}
	}
}

func StartWorker(ctx context.Context, name string, ch chan int) {
	w := &worker{
		ctx:  ctx,
		name: name,
		ch:   ch,
	}
	w.run()
	workers[name] = w
}
