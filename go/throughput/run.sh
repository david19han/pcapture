#!/bin/sh
sudo tcpdump -l -n -i en0 tcp | ./stdin
