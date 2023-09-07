package main

import "fmt"

func main() {
    fmt.Println("Customer database")

    customers := GetCustomers()

    for _, customer := range customers {
        fmt.Println(customer)
    }
}
