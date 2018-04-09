package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "time"
  "sync"
  // "sync/atomic"
  "strings"
  "strconv"
)

func getInput(input chan string) {
    for {
        in := bufio.NewReader(os.Stdin)
        result, err := in.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }
        input <- result
    }
}

func main() {
	input := make(chan string)
	state := make(map[string]int)
	m := &sync.Mutex{}

    go getInput(input)

    for {
        select {
        case text := <-input:
			if len(text) > 1{
				// fmt.Println(text)
				s := strings.Split(text," ")
				// fmt.Println(s)
				if(len(s) > 4){
					src := s[2]
					dst := s[4]
					// header := s[len(s)-2]
					length := s[len(s)-1]
					combined := src + " > " + dst //+ " > " + header + " > " + length
					pLength := length[0:len(length)-1]
					intVal,_ := strconv.Atoi(pLength)
					m.Lock()
					state[combined] = intVal
					m.Unlock()
					//atomic add to packet variable
					// atomic.AddUint32(&packet, uint32(intVal))
				}
			}else{
				fmt.Println("Too bad")
			}
        case <-time.After(1000 * time.Millisecond):
            start := time.Now()
			fmt.Println("One Second")
			//print out map
			total := 0
			m.Lock()
			for key, value := range state {
				total += value
			    fmt.Println("Key:", key, "Value:", value)
			}
			//flush map
			state = make(map[string]int)
			m.Unlock()
			//get current time again
      		t := time.Now()
			//print the elapsed time
      		elapsed := t.Sub(start)
      		fmt.Println("Total Length:",total,"|",elapsed,"elapsed")
        }
    }
}