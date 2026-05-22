// package main

// import (
//     "fmt"
//     "time"
// )

// func processTask(id int) {
//     fmt.Printf("Processing task %d...\n", id)
// }

// func main() {
//     // Spin up 3 separate tasks running concurrently
//     for i := 1; i <= 3; i++ {
//         fmt.Println("What works here?")
//         processTask(i) 
//     }

//     // Give the background tasks a brief second to execute before main exits
//     time.Sleep(100 * time.Millisecond)
// }


package main
import(
    "fmt"
    "time"
)

func processTask(i int )[
    fmt.Printf("The cpu is runnign task %d", i)
]


func main(){

    for i:= ; i<=3 ; i++{
        fmt.Println("what works here btw?")
        processTask(i)
    }

    time.Sleep(time.Millisecond *100)
}


// the goroutine can be chaotic, it foesnt guarantee that stuff would work th way you want
// The waitgroup and channel