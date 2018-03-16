package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "time"
  // "sync"
  // "strings"
  // "strconv" 
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
	// state := make(map[string]int)
	// packet := 0
	// m := &sync.Mutex{}

	ticker := time.NewTicker(time.Second).C

	for{
		select{
		case <- ticker:
			//print out map
			//print out packet variable
			//flush map 
			//set packet to 0 
			fmt.Println("One Second")
		// default: 
		// 	m.Lock()
		// 	text, err := reader.ReadString('\n')
		// 	m.Unlock()
		// 	if err == nil{
		// 		if len(text) > 1{
		// 			// fmt.Println(text)
		// 			s := strings.Split(text," ")
		// 			// fmt.Println(s)

		// 			src := s[2]
		// 			dst := s[4]
		// 			length := s[len(s)-1]

		// 			combined := src + " > " + dst
		// 			intVal,_ := strconv.Atoi(length) 
		// 			m.Lock()
		// 			state[combined] = intVal
		// 			m.Unlock()
		// 			//atomic add to packet variable 
		// 		}else{
		// 			fmt.Println("Too bad")
		// 		}
		// 	}else{
		// 		fmt.Println("Error")
		// 	}

		}
	}
}