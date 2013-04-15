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
