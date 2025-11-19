package main

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func main() {
	var free, total uint64
	path := "C:\\"
	windows.GetDiskFreeSpaceEx(path, &free, &total, nil)
	percentFree := float64(free) / float64(total) * 100
	if percentFree < 10 {
		fmt.Printf("Alerte ! Espace disque faible : %.2f%%\n", percentFree)
	} else {
		fmt.Printf("Espace disque OK : %.2f%%\n", percentFree)
	}
}
