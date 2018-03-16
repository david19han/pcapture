#!/bin/sh
GOOS=linux GOARCH=arm GOARM=5 go build $1.go
