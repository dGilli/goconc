package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
    start := time.Now()
    userName := fetchUser()
    respch := make(chan any, 2)
    wg := sync.WaitGroup{}

    wg.Add(2)
    go fetchUserLikes(userName, respch, &wg)
    go fetchUserPosts(userName, respch, &wg)

    wg.Wait() // block until 2 wg.Done() calls
    close(respch)

    for resp := range respch {
        fmt.Println("Response:", resp)
    }

    //fmt.Println("Likes:", likes)
    //fmt.Println("Posts:", posts)
    fmt.Println("Time taken:", time.Since(start))
}

func fetchUser() string {
    time.Sleep(time.Millisecond * 100)

    return "JOE"
}

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
    time.Sleep(time.Millisecond * 100)

    respch <- 42
    wg.Done()
}

func fetchUserPosts(userName string, respch chan any, wg *sync.WaitGroup) {
    time.Sleep(time.Millisecond * 150)

    respch <- "Post 1"
    wg.Done()
}
