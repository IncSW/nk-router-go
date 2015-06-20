package main

import (
    "fmt"
    "github.com/IncSW/nk-router-go"
)

func main() {
    router := nk.NewRouter()
    router.Add("main", nk.NewRoute("GET", "/", func() {}))
    router.Add("post", nk.NewRoute("GET", "/topic/:topic/post/:post", func() {}))
    
    if r, _ := router.Match("GET", "/"); r == nil {
        fmt.Println("404 Not Found")
    } else {
        fmt.Println("GET /")
    }
    
    if r, _ := router.Match("POST", "/"); r == nil {
        fmt.Println("404 Not Found")
    } else {
        fmt.Println("POST /")
    }

    if r, p := router.Match("GET", "/topic/12/post/34"); r == nil {
        fmt.Println("404 Not Found")
    } else {
        fmt.Println("GET /topic/12/post/34", p)
    }
}