package main

import (
	"arq"
	"arq/sr"
	"fmt"
	"time"
)

func main() {
	senderOut := make(chan arq.Packet)
	senderIn := make(chan arq.Packet)

	sender := sr.NewComputer(8, senderIn, senderOut)
	receiver := sr.NewComputer(8, senderOut, senderIn)

	go func() {
		if sequenceNumber, err := sender.Send(); err != nil {
			fmt.Println("Error - ", err)
		} else {
			fmt.Printf("Sender sent packet with sequence number %d\n", sequenceNumber)
		}
	}()

	go receiveHandler(receiver, "Receiver")
	go receiveHandler(sender, "Sender")

	time.Sleep(30 * time.Second)
}

func receiveHandler(c *sr.Computer, name string) {
	for {
		packet, err := c.Receive()

		if err != nil {
			fmt.Println("Error - ", err)
		} else {
			fmt.Printf("%s received: %v\n", name, packet)
		}
	}
}
