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

	go func() {
		for {
			packet, err := receiver.Receive()

			if err != nil {
				fmt.Println("Error - ", err)
			} else {
				fmt.Printf("Receiver received: %v\n", packet)
			}
		}
	}()

	go func() {
		for {
			packet, err := sender.Receive()

			if err != nil {
				fmt.Println("Error - ", err)
			} else {
				fmt.Printf("Sender received: %v\n", packet)
			}
		}
	}()

	time.Sleep(30 * time.Second)
}
