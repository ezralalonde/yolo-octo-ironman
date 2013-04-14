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
    var kk, nn uint
    for ii := 0; ii < lines; ii++ {
        _, err := fmt.Scanf("%d %d", &nn, &kk)
        if err != nil {
            panic("could not read n or k")
        }
        fmt.Printf("The bit patterns are\n")
        out := pat.GeneratePatterns(nn, kk)
        for _, oo := range out {
            fmt.Printf("%s\n", oo)
        }
        fmt.Printf("\n")
    }
}
