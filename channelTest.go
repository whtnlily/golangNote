package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)
	go handleMsg(ch)
	for index := range ch {
		fmt.Println(index)
	}
}


func handleMsg(ch chan int) {
	t := time.Tick(time.Second)
	index := 0
	for {
		select {
		case <-t:
			ch <- index
			index++
		}
	}
}