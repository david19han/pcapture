#!/bin/sh
sudo tcpdump -n -i en0 tcp | ./stdin
