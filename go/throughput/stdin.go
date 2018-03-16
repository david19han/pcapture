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

func hello() {
	fmt.Println("Hello world goroutine")
}

func main() {
	fmt.Println("Throughput Tool")
	fmt.Println("---------------------")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	_, err:= reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	state := make(map[string]int)
	// var packet uint32
	// packet = 0
	m := &sync.Mutex{}

	ticker := time.NewTicker(time.Second).C

	for{
		select{
		case <- ticker:
			
			//get current time
			fmt.Println("One Second")
			//print out map
			total := 0
			m.Lock()
			for key, value := range state {
				total += value 
			    fmt.Println("Key:", key, "Value:", value)
			}
			//print out packet variable
			fmt.Println("Total Length:",total)
			//flush map 
			state = make(map[string]int)
			//set packet to 0 
			m.Unlock()

			//get current time again
			//print the elapsed time
			
		default: 
			m.Lock()
			text, err := reader.ReadString('\n')
			m.Unlock()
			// if err == nil{
			// 	fmt.Println(text)
			// }
			
			if err == nil{
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
			}else{
				fmt.Println("Error")
			}

		}
	}
}