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
### Example 2
Send four packets from a sender that has a window size of two. The first packet gets lost in the network when sent from the sender. RTT in the network takes one second.
```
$ main -packet-sequence="S___" -window-size=2 -rtt=1s
14:49:07.621543 Sender sent packet with sequence number 0
14:49:07.872691 Sender sent packet with sequence number 1
14:49:08.874062 Receiver received: Packet #1
14:49:09.875348 Sender received: Packet #0 - ACK for 1
14:49:12.622697 Sender timeout triggered for Packet #0, resending...
14:49:13.623912 Receiver received: Packet #0
14:49:14.625071 Sender received: Packet #0 - ACK for 0
14:49:14.625142 Sender sent packet with sequence number 2
14:49:15.626331 Receiver received: Packet #2
14:49:16.627446 Sender received: Packet #0 - ACK for 2
14:49:16.627490 Sender sent packet with sequence number 3
14:49:17.628695 Receiver received: Packet #3
14:49:18.629983 Sender received: Packet #0 - ACK for 3
```
### Example 3
Send four packets from a sender that has a window size of two. The first packet gets lost in the network when sent from the sender along with the ACK for the packet. RTT in the network takes one second.
```
$ main -packet-sequence="B___" -window-size=2 -rtt=1s
16:01:31.701331 Sender sent packet with sequence number 0
16:01:31.952502 Sender sent packet with sequence number 1
16:01:32.953850 Receiver received: Packet #1
16:01:33.955104 Sender received: Packet #0 - ACK for 1
16:01:36.702485 Sender timeout triggered for Packet #0, resending...
16:01:37.703753 Receiver received: Packet #0
16:01:41.703103 Sender timeout triggered for Packet #0, resending...
16:01:42.704292 Receiver received: Packet #0
16:01:43.705595 Sender received: Packet #0 - ACK for 0
16:01:43.705766 Sender sent packet with sequence number 2
16:01:44.706295 Receiver received: Packet #2
16:01:45.706785 Sender received: Packet #0 - ACK for 2
16:01:45.706838 Sender sent packet with sequence number 3
16:01:46.707709 Receiver received: Packet #3
16:01:47.708932 Sender received: Packet #0 - ACK for 3
```