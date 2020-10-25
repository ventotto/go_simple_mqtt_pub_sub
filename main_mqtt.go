package main

import (
	connector "./internal"
	"fmt"
	"time"
)



func main() {
	topic := "test"

	go connector.Listen(topic)

	timer := time.NewTicker(1 * time.Second)
	noMsg := 1;
	for t := range timer.C {
		connector.Publisher(topic, fmt.Sprint("message_", noMsg, " (", t.String(), ")"))
		noMsg++;
	}
}
