package main

import (
        "fmt"
        "time"
)

func main() {
    testDates()
}

func testDates() {
    date := time.Now()

    fmt.Println("today is: ", date)

    customDate := time.Date(2023, 9, 6, 0, 0, 0, 0, time.UTC)

    fmt.Println("Custom time is: ", customDate)

    fmt.Println("Custom time year is: ", customDate.Year())
    fmt.Println("Custom time year is: ", customDate.Month())
    fmt.Println("Custom time year is: ", customDate.Day())

}
