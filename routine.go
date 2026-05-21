package main

import (
    "fmt"
    "time"
)

func processTask(id int) {
    fmt.Printf("Processing task %d...\n", id)
}

func main() {
    // Spin up 3 separate tasks running concurrently
    for i := 1; i <= 3; i++ {
        go processTask(i) 
    }

    // Give the background tasks a brief second to execute before main exits
    time.Sleep(100 * time.Millisecond)
}