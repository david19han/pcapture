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

    startTime := 0.01
    endTime := 0.01

    start_txt := ""
    end_txt := ""

    isStart := false 
    // isEnd := false

    go getInput(input)
    go timeout(timeChannel)

    for {
        select {
        case msg1 := <-timeChannel:
            start := time.Now()
            // fmt.Println("----------------------------------------")
            // fmt.Println(msg1)
            // fmt.Println("----------------------------------------")

            fmt.Println("----------------------------------------")
            fmt.Println()
            
            fmt.Println(msg1)
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

            //reset isStart
            isStart = false
            // isEnd = false
        
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
            fmt.Println("Real elapsed time is: ", endTime-startTime)
            fmt.Println("Start",start_txt,"| end",end_txt)
            fmt.Println("----------------------------------------")



        default:
            select{
                case text := <-input:
                    s := strings.Split(text," ")
                    if(len(s) > 4 && strings.Compare(s[1],"IP") == 0){

                        t_timeStamp := s[0]
                        timeStamp:= strings.Split(t_timeStamp,":")
                        secStr := strings.Split(timeStamp[len(timeStamp)-1],".")
                        sec,err1 :=strconv.Atoi(secStr[0])
                        dec,err2 :=strconv.Atoi(secStr[1])
                        if (err1 != nil) && (err2 != nil){
                            log.Fatal("bad conversion")
                        }

                        if !isStart {
                            startTime = float64(sec)+float64(dec)*math.Pow10(-1*len(secStr[1]))
                            start_txt = s[0]
                            isStart = true
                        }

                        if isStart{
                            endTime = float64(sec)+float64(dec)*math.Pow10(-1*len(secStr[1]))
                            end_txt = s[0]
                        }

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
                default:

            }       
        }
    }
}