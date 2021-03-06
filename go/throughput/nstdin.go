package main
 
import (
    "os"
    "bufio"
    "fmt"
    "time"
    "strings"
    "strconv"
    "log"
    "math"
)
 

// func main() {
//     scanner := bufio.NewScanner(os.Stdin)
//     for scanner.Scan() {
//         fmt.Println(scanner.Text())
//     }
//     if err := scanner.Err(); err != nil {
//         log.Println(err)
//     }
// }

func getInput(input chan string) {
    scanner := bufio.NewScanner(os.Stdin)
    // s := ""
    for scanner.Scan() {
        s := scanner.Text()
        // fmt.Println(scanner.Text())
        input <- s 
    }
}

func timeout(input chan string){
    for{
        // time.After(time.Second)
        time.Sleep(time.Second)
        input <- "One Second"
    }
}

func main() {
    input := make(chan string)
    timeChannel := make(chan string)

    state := make(map[string]int)
    // history := make(map[string][]string)

    go getInput(input)
    go timeout(timeChannel)

    // start_clock := time.Now()
    // end_clock := time.Now()

    for {
        select {
        case <-timeChannel:
            // end_clock = time.Now()
            start := time.Now()
            // fmt.Println("----------------------------------------")
            // fmt.Println()
            
            // fmt.Println(msg1)
            //print out map
            total := 0
            
            for _, value := range state {
                total += value
                // fmt.Println("Key:", key, "Value:", value)
            }
            // fmt.Println()

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
            // fmt.Println()
            // fmt.Println("Total Length (Bytes):",total,"|","Total Length (MB):",total_mb,"|",
            //     "Total time:",diff,"|", "TPut:", total_mb/diff_out)
            fmt.Println(total_mb/diff_out)

            // total_elapsed := float64(end_clock.Sub(start_clock) - elapsed)* math.Pow10(-9)
            // fmt.Println("True elapsed time:",end_clock.Sub(start_clock),elapsed,total_elapsed)
            // fmt.Println("True Tput:",total_mb/total_elapsed)
            // start_clock = time.Now()

            // fmt.Println("----------------------------------------")

            


        default:
            select{
                case text := <-input:
                    s := strings.Split(text," ")
                    // bool b1 := (strings.Compare(s[1],"IP") == 0)
                    // bool b2 := (strings.Compare(s[1],"IP6") == 0)
                    if(len(s) > 4 && ( (strings.Compare(s[1],"IP") == 0) || (strings.Compare(s[1],"IP6") == 0))){
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

                        // fmt.Println("Adding:",combined,"| Packets:",intVal,"| Text:")
                        // fmt.Println(text)
                        // fmt.Println()
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
                default:

            }       
        }
    }
}