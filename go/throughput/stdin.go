package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "strings"
  // "time"
  // "sync" 
)

func main() {
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	reader := bufio.NewReader(os.Stdin)
	m := &sync.Mutex{}

	fmt.Print("-> ")
	_, err:= reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	m := make(map[string]int)
	packet := 0
	for{
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("Total Capture =",packet)
			for key, value := range m {
   				fmt.Println("Key:", key, "Value:", value)
			} 
			log.Fatal(err)
		}else{
			if len(text) > 1{
				fmt.Println(text)
				s := strings.Split(text," ")

				src := s[2]
				dst := s[4]
				length := s[len(s)-1]

				combined := src + " > " + dst + " > " + length
				//lock
				v, prs := m[combined]
				if prs{
					m[combined] = v+1
				}else{
					m[combined] = 1
				}
				//unlock
				fmt.Println(combined)
				
			}
			packet++
		}
	}
  

}
