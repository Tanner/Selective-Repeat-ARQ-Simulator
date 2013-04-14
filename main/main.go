package main

import (
	"arq"
	"arq/sr"
	"log"
	"time"
)

func main() {
	senderOut := make(chan arq.Packet)
	senderIn := make(chan arq.Packet)

	sender := sr.NewComputer(8, senderIn, senderOut)
	receiver := sr.NewComputer(8, senderOut, senderIn)

	for i := 0; i < 9; i++ {
		go send(sender)
	}

	go receiveHandler(sender, "Sender")
	go receiveHandler(receiver, "Receiver")

	time.Sleep(30 * time.Second)
}

func send(c *sr.Computer) {
	if sequenceNumber, err := c.Send(); err != nil {
		log.Println("Error - ", err)
	} else {
		log.Printf("Sender sent packet with sequence number %d\n", sequenceNumber)
	}
}

func receiveHandler(c *sr.Computer, name string) {
	for {
		packet, err := c.Receive()

		if err != nil {
			log.Println("Error - ", err)
		} else {
			log.Printf("%s received: %v\n", name, packet)
		}
	}
}
