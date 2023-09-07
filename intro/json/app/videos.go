package main

import (
    "io/ioutil"
    "encoding/json"
)

type video struct {
        Id string `json:"id"`
        Title string `json:"title"`
        Description string `json:"description"`
        Url string `json:"url"`
}

func getVideos() (videos []video) {
    filebytes, err := ioutil.ReadFile("./videos.json")
    
    if err != nil {
            panic(err)
    }

    err = json.Unmarshal(filebytes, &videos)

    if err != nil {
            panic(err)
    }

    return videos
}
