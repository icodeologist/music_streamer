package main

import "time"

func main() {
	go StartandListenServer()
	time.Sleep(2 * time.Second)
	ClientStart()
}
