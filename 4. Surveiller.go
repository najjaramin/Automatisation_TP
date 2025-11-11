package main

import (
    "fmt"
    "syscall"
)

func main() {
    var stat syscall.Statfs_t
    err := syscall.Statfs("/", &stat)
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }

    total := stat.Blocks * uint64(stat.Bsize)
    free := stat.Bavail * uint64(stat.Bsize)
    percentFree := float64(free) / float64(total) * 100

    if percentFree < 10 {
        fmt.Println("⚠️ Alerte : espace disque libre inférieur à 10%")
    } else {
        fmt.Printf("Espace libre : %.2f%%\n", percentFree)
    }
}