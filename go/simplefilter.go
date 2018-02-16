package main

import (
    "fmt"
    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "log"
    "time"
)

var (
    device       string = "en0"
    snapshot_len int32  = 1024
    promiscuous  bool   = false
    err          error
    timeout      time.Duration = 30 * time.Second
    handle       *pcap.Handle
)

func main() {
    // Open device
    handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
    if err != nil {
        log.Fatal(err)
    }
    defer handle.Close()

    // Set filter
    var filter string = "tcp" // include udp 
    err = handle.SetBPFFilter(filter)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Only capturing tcp")

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

    tPut := make(map[gopacket.Endpoint]int)
    i := 0
    for packet := range packetSource.Packets() {
        netFlow := packet.NetworkLayer().NetworkFlow()
        src,dst := netFlow.Endpoints()
        fmt.Println("src:",src,"dst:",dst)//4 tuple including src port and dst port

        if val, ok := tPut[src]; ok {
            tPut[src] = val + 1
        }else{
            tPut[src] = 1
        }

        i++
        if(i > 5){
            break
        }
    }

    //print table
    for k,v := range tPut{
        fmt.Println("Key:", k, "Value:", v)
    }

}