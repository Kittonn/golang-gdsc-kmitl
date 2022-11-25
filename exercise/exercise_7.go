package exercise

import (
	"context"
	"fmt"
)

func send(ch chan<- int, cancel context.CancelFunc) {
	defer cancel()
	for i := 0; i < 100; i++ {
		ch <- 1
	}
}

func receive(ch <-chan int, ctx context.Context) {
	var datas int
	for {
		select {
		case data := <-ch:
			datas += data
		case <-ctx.Done():
			fmt.Println(datas)
			return
		}
	}
}

func Exercise_7() {
	ch := make(chan int)
	bgCTX := context.Background()
	ctx, cancel := context.WithCancel(bgCTX)
	go send(ch, cancel)
	receive(ch, ctx)
}
