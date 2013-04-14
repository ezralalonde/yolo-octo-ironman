package main

import (
    "fmt"
    "yolo-octo-ironman/pat"
)

func main() {
    var lines int
    _, err := fmt.Scanf("%d", &lines)
    if err != nil {
        panic("could not read number of lines")
    }
    var kk, nn int
    for ii := 0; ii < lines; ii++ {
        _, err := fmt.Scanf("%d %d", &nn, &kk)
        if err != nil {
            panic("could not read n or k")
        }
        fmt.Printf("The bit patterns are\n")
        bits := pat.GeneratePatterns(nn, kk)
        for _, bit := range bits {
            fmt.Println(bit)
        }
        fmt.Printf("\n")
    }
}
