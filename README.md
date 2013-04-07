CS 4251 - Programming Project
-----------------------------
You need to write a program to emulate the sliding window behavior of a sender process that uses selective repeat ARQ with a window size of 8 packets. Your program should allow users to specify the number of packets to be "sent" to a receiver process and which subset of them will be "lost in the network". 

In this emulation, you have the liberty to specify the roundtrip delay and the times these packets arrive at the sender process, ready to be sent out. 

You need NOT write any TCP/UDP socket programs. Instead, you may simply assume that each packet that is not "lost in the network" will magically be received and the sender process will magically get a corresponding acknowledgement from the receiver. Then, the timers associated with these lost packets will expire after a certain amount of time (say 5 seconds) and they will be "resent".

Your program needs to output which packet (identified by a sequence number) will be "sent" out at what time.