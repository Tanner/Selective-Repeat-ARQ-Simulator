CS 4251 - Programming Project
-----------------------------
A simulator of the sliding window behavior of a sender process that uses select repeat ARQ.

Usage
-----
```
Usage of main:
  -packet-sequence="__S_": The sequence of packets to send. '_' no losses, 'A' ACK loss, 'S', sender loss, 'B' both lost
  -packet-time=250ms: Amount of time waited after each packet is sent.
  -rtt=200ms: Round trip time between a packet being sent and the acknowledgment returning.
  -timeout=5s: Amount of time to wait before resending a packet that hasn't been acknowledged.
  -window-size=8: Window size for the selective repeat protocol
```