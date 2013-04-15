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
	timeBetweenPackets := flag.Duration("packet-time", 250*time.Millisecond, "Amount of time waited after each packet is sent.")
	timeoutDuration := flag.Duration("timeout", 5*time.Second, "Amount of time to wait before resending a packet that hasn't been acknowledged.")
	roundTripDuration := flag.Duration("rtt", 200*time.Millisecond, "Round trip time between a packet being sent and the acknowledgment returning.")
	windowSize := flag.Int("window-size", 8, "Window size for the selective repeat protocol")

	flag.Parse()

	packetLoss := parsePacketSequence(*packetSequence)

	senderOut := make(chan arq.Packet)
	senderIn := make(chan arq.Packet)

	receivedACK := make(chan int)

	sender := sr.NewComputer(*windowSize, senderIn, senderOut, *roundTripDuration, *timeoutDuration, senderTimeoutTriggered)
	receiver := sr.NewComputer(*windowSize, senderOut, senderIn, *roundTripDuration, *timeoutDuration, nil)

	go receiveHandler(sender, "Sender", receivedACK)
	go receiveHandler(receiver, "Receiver", nil)

	go func() {
		for _, v := range packetLoss {
			go send(sender, v.Sender, v.Acknowledgment)

			time.Sleep(*timeBetweenPackets)
		}
	}()

	for _ = range packetLoss {
		<-receivedACK
	}
}

// parsePacketSequence returns an array of PacketLoss for each char in the packet sequence
func parsePacketSequence(packetSequence string) []PacketLoss {
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

// send sends a packet from the given computer
// Values of senderLoss and acknowledgementLoss impact whether the packet is lost at the sender level or the receiver ACK level
func send(c *sr.Computer, senderLoss, acknowledgementLoss bool) {
	if sequenceNumber, err := c.Send(senderLoss, acknowledgementLoss); err != nil {
		log.Println("Error - ", err)
	} else {
		log.Printf("Sender sent packet with sequence number %d\n", sequenceNumber)
	}
}

// receiveHandler receives packets on the computer and passes knowledge of any ACKS to the channel
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

// senderTimeoutTriggered logs when a timeout has occurred
func senderTimeoutTriggered(sequenceNumber int) {
	log.Printf("Sender timeout triggered for Packet #%d, resending...", sequenceNumber)
}
