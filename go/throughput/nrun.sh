#!/bin/sh
sudo tcpdump -l -n -i en0 tcp or udp | ./nStdin 
