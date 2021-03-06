package main

import (
  "bufio"
  "fmt"
  "os"
  "log"
  "time"
  // "sync"
  // "sync/atomic"
  "strings"
  "strconv"
  "math"
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
    // history := make(map[string][]string)

    go getInput(input)

    for {
        select {
        case text := <-input:
            text = text[0:len(text)-1]
			s := strings.Split(text," ")
            if(len(s) > 4 && strings.Compare(s[1],"IP") == 0){
                src := s[2]
                dst := s[4]
                combined := src + " > " + dst 

                i := 0
                lfound := false 
                for i = range s{
                    if strings.Compare(s[i],"length") == 0{
                        lfound = true 
                        break
                    }
                }

                intVal := 0 

                if(i+1 > len(s) -1){
                    last := s[len(s)-1]
                    if last[0] == '(' {
                        intVal,_ = strconv.Atoi(last[1:len(last)-1])
                    }else{
                        fmt.Println("ERROR")
                        fmt.Println(text)
                        fmt.Println(s)
                        log.Fatal("index over")
                    }
                }else {
                    if lfound {
                        intVal,_ = strconv.Atoi(s[i+1])
                        }else{
                            fmt.Println("ERROR")
                            fmt.Println(text)
                            fmt.Println(s)
                            log.Fatal("length field and lparenth not found")
                        }
                    
                }

                

                _,state_ok := state[combined]
                if state_ok{
                    state[combined] += intVal
                }else{
                    state[combined] = intVal
                }

                // history[combined] = append(history[combined], text)

                fmt.Println("Adding:",combined,"| Packets:",intVal,"| Text:")
                fmt.Println(text)
                fmt.Println()
            }else{
                fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
                fmt.Println()
                fmt.Println("ERROR")
                fmt.Println(text)
                fmt.Println(s)
                fmt.Println("IP is not the second field")
                fmt.Println()
                fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
            }
        case <-time.After(1000 * time.Millisecond):
            start := time.Now()
            // fmt.Println("----------------------------------------")
            // fmt.Println(msg1)
            // fmt.Println("----------------------------------------")

            fmt.Println("----------------------------------------")
            fmt.Println()
            
            fmt.Println("One Second")
            //print out map
            total := 0
            
            for key, value := range state {
                total += value
                fmt.Println("Key:", key, "Value:", value)
            }
            fmt.Println()
            // fmt.Println("History")
            // for key, value := range history{
            //     fmt.Println("-Key:", key, "PACKETS:")
            //     fmt.Printf("%v", value)
            //     fmt.Println()
            // }

            //flush map
            state = make(map[string]int)
            // history = make(map[string][]string)
        
            //get current time again
            t := time.Now()
            //print the elapsed time
            elapsed := t.Sub(start)
            diff := time.Second - elapsed 

            total_mb := float64(total)*8*math.Pow10(-6)
            diff_out := float64(diff) * math.Pow10(-9)
            fmt.Println()
            fmt.Println("Total Length (Bytes):",total,"|","Total Length (MB):",total_mb,"|",
                "Total time:",diff,"|", "TPut:", total_mb/diff_out)
            fmt.Println("----------------------------------------")
        }
    }
}