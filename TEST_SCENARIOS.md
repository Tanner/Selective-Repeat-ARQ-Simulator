## Test Scenarios
### Scenario 1
```
$ main -packet-time=0s -rtt=1s -packet-sequence="_S_________" -window-size=8
22:15:18.383773 Sender sent packet with sequence number 0
22:15:18.383936 Sender sent packet with sequence number 1
22:15:18.383953 Sender sent packet with sequence number 2
22:15:18.383998 Sender sent packet with sequence number 3
22:15:18.384006 Sender sent packet with sequence number 4
22:15:18.384014 Sender sent packet with sequence number 5
22:15:18.384022 Sender sent packet with sequence number 6
22:15:18.384038 Sender sent packet with sequence number 7
22:15:18.885208 Receiver received: Packet #0
22:15:18.885328 Receiver received: Packet #2
22:15:18.885362 Receiver received: Packet #3
22:15:18.885376 Receiver received: Packet #4
22:15:18.885408 Receiver received: Packet #7
22:15:18.885451 Receiver received: Packet #5
22:15:18.885485 Receiver received: Packet #6
22:15:19.386316 Sender received: Packet #0 - ACK for 0
22:15:19.386413 Sender received: Packet #0 - ACK for 2
22:15:19.386445 Sender received: Packet #0 - ACK for 3
22:15:19.386403 Sender sent packet with sequence number 8
22:15:19.386489 Sender received: Packet #0 - ACK for 4
22:15:19.386505 Sender received: Packet #0 - ACK for 7
22:15:19.386530 Sender received: Packet #0 - ACK for 5
22:15:19.386547 Sender received: Packet #0 - ACK for 6
22:15:19.887744 Receiver received: Packet #8
22:15:20.388895 Sender received: Packet #0 - ACK for 8
22:15:23.384765 Sender timeout triggered for Packet #1, resending...
22:15:23.885965 Receiver received: Packet #1
22:15:24.387096 Sender received: Packet #0 - ACK for 1
22:15:24.387154 Sender sent packet with sequence number 9
22:15:24.387168 Sender sent packet with sequence number 10
22:15:24.888318 Receiver received: Packet #9
22:15:24.888399 Receiver received: Packet #10
22:15:25.389180 Sender received: Packet #0 - ACK for 9
22:15:25.389220 Sender received: Packet #0 - ACK for 10
```

### Scenario 2
```
$ main -packet-time=0s -rtt=1s -packet-sequence="SSS________" -window-size=8
19:36:14.416601 Sender sent packet with sequence number 0
19:36:14.416789 Sender sent packet with sequence number 1
19:36:14.416801 Sender sent packet with sequence number 2
19:36:14.416815 Sender sent packet with sequence number 3
19:36:14.416823 Sender sent packet with sequence number 4
19:36:14.416830 Sender sent packet with sequence number 5
19:36:14.416838 Sender sent packet with sequence number 6
19:36:14.416851 Sender sent packet with sequence number 7
19:36:14.918216 Receiver received: Packet #3
19:36:14.918249 Receiver received: Packet #4
19:36:14.918269 Receiver received: Packet #5
19:36:14.918313 Receiver received: Packet #6
19:36:14.918334 Receiver received: Packet #7
19:36:15.419419 Sender received: Packet #0 - ACK for 3
19:36:15.419553 Sender received: Packet #0 - ACK for 4
19:36:15.419660 Sender received: Packet #0 - ACK for 5
19:36:15.419728 Sender received: Packet #0 - ACK for 6
19:36:15.419787 Sender received: Packet #0 - ACK for 7
19:36:19.417777 Sender timeout triggered for Packet #0, resending...
19:36:19.417882 Sender timeout triggered for Packet #1, resending...
19:36:19.417889 Sender timeout triggered for Packet #2, resending...
19:36:19.918112 Receiver received: Packet #0
19:36:19.918155 Receiver received: Packet #1
19:36:19.918184 Receiver received: Packet #2
19:36:20.419237 Sender received: Packet #0 - ACK for 0
19:36:20.419299 Sender received: Packet #0 - ACK for 1
19:36:20.419332 Sender received: Packet #0 - ACK for 2
19:36:20.419322 Sender sent packet with sequence number 9
19:36:20.419268 Sender sent packet with sequence number 8
19:36:20.419354 Sender sent packet with sequence number 10
19:36:20.920409 Receiver received: Packet #8
19:36:20.920458 Receiver received: Packet #9
19:36:20.920495 Receiver received: Packet #10
19:36:21.421605 Sender received: Packet #0 - ACK for 8
19:36:21.421645 Sender received: Packet #0 - ACK for 9
19:36:21.421681 Sender received: Packet #0 - ACK for 10
```

### Scenario 3
```
$ main -packet-time=0s -rtt=1s -packet-sequence="SAB________" -window-size=8
22:28:06.716283 Sender sent packet with sequence number 0
22:28:06.716543 Sender sent packet with sequence number 1
22:28:06.716550 Sender sent packet with sequence number 2
22:28:06.716560 Sender sent packet with sequence number 3
22:28:06.716567 Sender sent packet with sequence number 4
22:28:06.716575 Sender sent packet with sequence number 5
22:28:06.716582 Sender sent packet with sequence number 6
22:28:06.716596 Sender sent packet with sequence number 7
22:28:07.217629 Receiver received: Packet #1
22:28:07.217674 Receiver received: Packet #3
22:28:07.217738 Receiver received: Packet #4
22:28:07.217829 Receiver received: Packet #5
22:28:07.217865 Receiver received: Packet #6
22:28:07.217904 Receiver received: Packet #7
22:28:07.718036 Sender received: Packet #0 - ACK for 3
22:28:07.718152 Sender received: Packet #0 - ACK for 4
22:28:07.718257 Sender received: Packet #0 - ACK for 5
22:28:07.718330 Sender received: Packet #0 - ACK for 6
22:28:07.718374 Sender received: Packet #0 - ACK for 7
22:28:11.716549 Sender timeout triggered for Packet #0, resending...
22:28:11.716572 Sender timeout triggered for Packet #1, resending...
22:28:11.716577 Sender timeout triggered for Packet #2, resending...
22:28:12.217701 Receiver received: Packet #0
22:28:12.217773 Receiver received: Packet #1
22:28:12.217826 Receiver received: Packet #2
22:28:12.718727 Sender received: Packet #0 - ACK for 0
22:28:12.718878 Sender received: Packet #0 - ACK for 1
22:28:12.718848 Sender sent packet with sequence number 8
22:28:12.718992 Sender sent packet with sequence number 9
22:28:13.219900 Receiver received: Packet #8
22:28:13.220029 Receiver received: Packet #9
22:28:13.720988 Sender received: Packet #0 - ACK for 8
22:28:13.721118 Sender received: Packet #0 - ACK for 9
22:28:16.717764 Sender timeout triggered for Packet #2, resending...
22:28:17.218915 Receiver received: Packet #2
22:28:17.719577 Sender received: Packet #0 - ACK for 2
22:28:17.719717 Sender sent packet with sequence number 10
22:28:18.220901 Receiver received: Packet #10
22:28:18.722121 Sender received: Packet #0 - ACK for 10
```