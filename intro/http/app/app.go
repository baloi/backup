package main

import (
    "net/http"
    "fmt"
    "os"
    "errors"
)

func main() {
    http.HandleFunc("/", Hello)
    err := http.ListenAndServe(":8080", nil)
    if errors.Is(err, http.ErrServerClosed) {
        fmt.Printf("server closed\n")
    } else if err != nil {
        fmt.Printf("error starting server %s \n", err)
        os.Exit(1)
    }
}

func Hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))   
}
