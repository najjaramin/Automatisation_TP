package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	url := "https://example.com"
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Erreur de connexion")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("Site accessible")
	} else {
		fmt.Println("Site inaccessible", resp.StatusCode)
	}
}
