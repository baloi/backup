package main

import "fmt"

func main() {
    videos := getVideos()

    for i, _ := range(videos) { 
        videos[i].Description = videos[i].Description + "\nRemember to Like and Subscribe!\n"
    }

    fmt.Println(videos)

    //@todo
    // saveVideos(videos)
}
