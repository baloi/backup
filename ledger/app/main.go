package main

import (
    "fmt"
)

func main() {
    var y string
    var x bool

    fmt.Println("--- List ---")

    for { 
        fmt.Println("Please enter choice")
        fmt.Println("[l] List")
        fmt.Println("[x] Exit")
        //fmt.Println("[a] Add")

        fmt.Scanln(&y)

        switch y {
            case "l":
                fmt.Println("----- Listing ------") 
                entries := getEntries()
                listEntries(entries)
                fmt.Println("--------------------") 
            case "x":
                fmt.Println("Exiting...")
                x = true 
        }
                
        if x {
            break
        }
    }
    fmt.Println("Have a good day!")
}
