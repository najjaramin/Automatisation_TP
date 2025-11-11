package main

import (
    "fmt"
    "os/exec"
)

func main() {
    out, err := exec.Command("who").Output()
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }
    fmt.Println("Utilisateurs connectés :\n", string(out))
}