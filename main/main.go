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
	timeBetweenPackets := flag.Duration("packet-time", 250000000, "Amount of time waited after each packet is sent.")
	timeoutDuration := flag.Duration("timeout", 5*time.Second, "Amount of time to wait before resending a packet that hasn't been acknowledged.")

	flag.Parse()

	packetLoss := parseArgs(*packetSequence)

	senderOut := make(chan arq.Packet)
	senderIn := make(chan arq.Packet)

	receivedACK := make(chan int)

	sender := sr.NewComputer(8, senderIn, senderOut, *timeoutDuration, senderTimeoutTriggered)
	receiver := sr.NewComputer(8, senderOut, senderIn, *timeoutDuration, nil)

	for _, v := range packetLoss {
		go send(sender, v.Sender, v.Acknowledgment)

		time.Sleep(*timeBetweenPackets)
	}

	go receiveHandler(sender, "Sender", receivedACK)
	go receiveHandler(receiver, "Receiver", nil)

	for _ = range packetLoss {
		<-receivedACK
	}
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

func receiveHandler(c *sr.Computer, name string, receivedACK chan int) {
	for {
		packet, err := c.Receive()

		if packet.ACK {
			receivedACK <- 1
		}

		if err != nil {
			log.Println("Error - ", err)
		} else {
			log.Printf("%s received: %v\n", name, packet)
		}
	}
}

func senderTimeoutTriggered(sequenceNumber int) {
	log.Printf("Sender timeout triggered for Packet #%d, resending...", sequenceNumber)
}
