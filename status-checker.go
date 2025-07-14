package main

import (
    "fmt"
    "net/http"
    "os"
    "sync"
    "time"
)

func checkStatus(url string, wg *sync.WaitGroup) {
    defer wg.Done()
    client := http.Client{
        Timeout: 3 * time.Second,
    }

    resp, err := client.Get(url)
    if err != nil {
        fmt.Printf("%s ❌ DOWN (%v)\n", url, err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("%s ✅ UP (%d %s)\n", url, resp.StatusCode, http.StatusText(resp.StatusCode))
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run statuschecker.go <url1> <url2> ...")
        return
    }

    var wg sync.WaitGroup
    urls := os.Args[1:]

    for _, url := range urls {
        wg.Add(1)
        go checkStatus(url, &wg)
    }

    wg.Wait()
}