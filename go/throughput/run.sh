#!/bin/sh
sudo tcpdump -n -c 10 -i en0 tcp | ./stdin
