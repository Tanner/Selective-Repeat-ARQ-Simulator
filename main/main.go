package main

import (
	"arq"
	"arq/sr"
	"flag"
	"log"
	"time"
)

type PacketLoss struct {
	Sender         bool
	Acknowledgment bool
}

func main() {
	log.SetFlags(log.Lmicroseconds)

	packetSequence := flag.String("packet-sequence", "__S_", "The sequence of packets to send. '_' no losses, 'A' ACK loss, 'S', sender loss, 'B' both lost")
	flag.Parse()

	packetLoss := parseArgs(*packetSequence)

	senderOut := make(chan arq.Packet)
	senderIn := make(chan arq.Packet)

	sender := sr.NewComputer(8, senderIn, senderOut)
	receiver := sr.NewComputer(8, senderOut, senderIn)

	for _, v := range packetLoss {
		go send(sender, v.Sender, v.Acknowledgment)
	}

	go receiveHandler(sender, "Sender")
	go receiveHandler(receiver, "Receiver")

	time.Sleep(30 * time.Second)
}

func parseArgs(packetSequence string) []PacketLoss {
	loss := make([]PacketLoss, len(packetSequence))

	for i, v := range packetSequence {
		packetLoss := PacketLoss{false, false}

		switch v {
		case 'A', 'a':
			packetLoss.Acknowledgment = true
		case 'S', 's':
			packetLoss.Sender = true
		case 'B', 'b':
			packetLoss.Sender = true
			packetLoss.Acknowledgment = true
		}

		loss[i] = packetLoss
	}

	return loss
}

func send(c *sr.Computer, senderLoss, acknowledgementLoss bool) {
	if sequenceNumber, err := c.Send(senderLoss, acknowledgementLoss); err != nil {
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
