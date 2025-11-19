package main

import (
    "fmt"
    "net/http"
)

func main() {
    resp, err := http.Get("https://example.com")
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }
    if resp.StatusCode == 200 {
        fmt.Println("Site accessible")
    } else {
        fmt.Println("Site inaccessible")
    }
}
