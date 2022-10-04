package Datatype

import (
	"fmt"
	"math/rand"
	"time"
)

type T = struct{}

func LontimedOperation() <-chan int32 {
	ch := make(chan int32)
	go func() {
		defer close(ch)
		time.Sleep(time.Second * 3)
		ch <- rand.Int31n(300)
	}()

	return ch
}

func Longtime() {
	completed := make(chan T)
	go func() {
		fmt.Println("ping")
		time.Sleep(time.Second * 5)
		<-completed
	}()
	completed <- struct{}{}
	fmt.Println("pong")
}
