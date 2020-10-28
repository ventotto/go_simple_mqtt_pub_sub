package main

import (
	connector "./internal"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	topic := "test"

	go connector.Listen(topic)

	select {}
}
