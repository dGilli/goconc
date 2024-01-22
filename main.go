package main

import (
	"fmt"
	"time"
)

func main() {
    start := time.Now()
    userName := fetchUser()
    likes := fetchUserLikes(userName)
    posts := fetchUserPosts(userName)

    fmt.Println("Likes:", likes)
    fmt.Println("Posts:", posts)
    fmt.Println("Time taken:", time.Since(start))
}

func fetchUser() string {
    time.Sleep(time.Millisecond * 100)

    return "JOE"
}

func fetchUserLikes(userName string) int {
    time.Sleep(time.Millisecond * 100)

    return 42
}

func fetchUserPosts(userName string) int {
    time.Sleep(time.Millisecond * 150)

    return 3
}
