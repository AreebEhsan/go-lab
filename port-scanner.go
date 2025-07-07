package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "sync"
    "time"
)

func scanPort(protocol, hostname string, port int, wg *sync.WaitGroup) {
    defer wg.Done()
    address := fmt.Sprintf("%s:%d", hostname, port)
    conn, err := net.DialTimeout(protocol, address, 1*time.Second)
    if err != nil {
        return // closed
    }
    conn.Close()
    fmt.Printf("Port %d is open\n", port)
}

func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run portscanner.go <host> <startPort> <endPort>")
        return
    }

    hostname := os.Args[1]
    startPort, err1 := strconv.Atoi(os.Args[2])
    endPort, err2 := strconv.Atoi(os.Args[3])

    if err1 != nil || err2 != nil || startPort < 1 || endPort < startPort {
        fmt.Println("Invalid port range")
        return
    }

    fmt.Printf("Scanning %s from port %d to %d...\n", hostname, startPort, endPort)

    var wg sync.WaitGroup
    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort("tcp", hostname, port, &wg)
    }

    wg.Wait()
    fmt.Println("Scan complete.")
}