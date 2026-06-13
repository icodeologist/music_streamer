package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go StartandListenServer()
	defer wg.Done()
	time.Sleep(2 * time.Second)
	ClientStart()
	wg.Wait()
}
