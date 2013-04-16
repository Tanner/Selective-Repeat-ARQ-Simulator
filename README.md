## CS 4251 - Programming Project

A simulator of the sliding window behavior of a sender process that uses select repeat ARQ.

## Usage
```
Usage of main:
  -packet-sequence="__S_": The sequence of packets to send. '_' no losses, 'A' ACK loss, 'S', sender loss, 'B' both lost
  -packet-time=250ms: Amount of time waited after each packet is sent.
  -rtt=200ms: Round trip time between a packet being sent and the acknowledgment returning.
  -timeout=5s: Amount of time to wait before resending a packet that hasn't been acknowledged.
  -window-size=8: Window size for the selective repeat protocol
```

### Packet Sequence
Packet sequence is a string of characters that indicate how many packets should be sent and when, if any at all, that packet or that packet's acknowledgement should be lost.

`_` indicates a packet that should be sent/received with no losses

`A` indicates a packet that should be successfully sent from the sender to the receiver, but the receiver's acknowledgement gets lost in the network

`S` indicates a packet that will not successfully sent from the sender to the receiver

`B` indicates a packet that will have both the behaviors of `A` and `B`. It will first fail to send the packet from the sender, resend, successfully arrive at the destination, the acknowledgement will get lost, and resend again with no further losses.

## Examples

### Example 1
Send four packets, without any losses, from a sender that has a window size of two. RTT in the network takes one second.
```
$ main -packet-sequence="____" -window-size=2 -rtt=1s

20:30:26.235480 Sender sent packet with sequence number 0
20:30:26.486683 Sender sent packet with sequence number 1
20:30:27.236893 Receiver received: Packet #0
20:30:27.487986 Receiver received: Packet #1
20:30:28.238078 Sender received: Packet #0 - ACK for 0
20:30:28.238138 Sender sent packet with sequence number 2
20:30:28.489267 Sender received: Packet #0 - ACK for 1
20:30:28.489387 Sender sent packet with sequence number 3
20:30:29.239264 Receiver received: Packet #2
20:30:29.490599 Receiver received: Packet #3
20:30:30.240425 Sender received: Packet #0 - ACK for 2
20:30:30.491180 Sender received: Packet #0 - ACK for 3
```
### Example 2
Send four packets from a sender that has a window size of two. The first packet gets lost in the network when sent from the sender. RTT in the network takes one second.
```
$ main -packet-sequence="S___" -window-size=2 -rtt=1s
22:16:19.856486 Sender sent packet with sequence number 0
22:16:20.107142 Sender sent packet with sequence number 1
22:16:20.607861 Receiver received: Packet #1
22:16:21.108971 Sender received: Packet #0 - ACK for 1
22:16:24.857674 Sender timeout triggered for Packet #0, resending...
22:16:25.358866 Receiver received: Packet #0
22:16:25.859544 Sender received: Packet #0 - ACK for 0
22:16:25.859603 Sender sent packet with sequence number 2
22:16:25.859615 Sender sent packet with sequence number 3
22:16:26.360736 Receiver received: Packet #2
22:16:26.360855 Receiver received: Packet #3
22:16:26.861968 Sender received: Packet #0 - ACK for 2
22:16:26.862100 Sender received: Packet #0 - ACK for 3
```
### Example 3
Send four packets from a sender that has a window size of two. The first packet gets lost in the network when sent from the sender along with the ACK for the packet. RTT in the network takes one second.
```
$ main -packet-sequence="B___" -window-size=2 -rtt=1s
22:16:46.744080 Sender sent packet with sequence number 0
22:16:46.994655 Sender sent packet with sequence number 1
22:16:47.495383 Receiver received: Packet #1
22:16:47.996357 Sender received: Packet #0 - ACK for 1
22:16:51.745324 Sender timeout triggered for Packet #0, resending...
22:16:52.246346 Receiver received: Packet #0
22:16:56.746606 Sender timeout triggered for Packet #0, resending...
22:16:57.247807 Receiver received: Packet #0
22:16:57.749014 Sender received: Packet #0 - ACK for 0
22:16:57.749106 Sender sent packet with sequence number 2
22:16:57.749122 Sender sent packet with sequence number 3
22:16:58.250238 Receiver received: Packet #2
22:16:58.250314 Receiver received: Packet #3
22:16:58.751251 Sender received: Packet #0 - ACK for 2
22:16:58.751386 Sender received: Packet #0 - ACK for 3
```