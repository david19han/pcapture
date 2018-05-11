#!/bin/sh
sudo tcpdump -l -n -i en0 udp or tcp | ./nStdin 
