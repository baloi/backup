package main

import (
    "os"
    "encoding/json"
    "fmt"
    "strconv"
)

type entry struct {
    Id string `json:"id"`
    Name string `json:"name"` 
    Description string `json:"description"` 
    DayOfMonth string `json:"dayofmonth"` 
    Kind string `json:"kind"` // expense or income
    Amount string `json:"amount"`
    //Id string `json:"id"`
    //Id string `json:"id"`
    /*
        theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
        fmt.Println("The time is", theTime)
    */
}

func getEntries() (entries []entry) {
    filebytes, err := os.ReadFile("./entries.json")

    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(filebytes, &entries)

    if err != nil {
        panic(err)
    }

    return entries
}

func listEntries(entries []entry) {
    for i, _ := range(entries) {
        fmt.Println("<" + strconv.Itoa(i) + ">")
        fmt.Println(entries[i].Name)
        fmt.Println(entries[i].Description)
        fmt.Println(entries[i].Kind)
        fmt.Println("$" + entries[i].Amount)
        fmt.Println("Day of month: " + entries[i].DayOfMonth)
    }
}
