package main

import (
    "fmt"
    "os"
    "sort"
)

type FileInfo struct {
    Name string
    Size int64
}

func main() {
    entries, _ := os.ReadDir(".")
    var files []FileInfo

    for _, entry := range entries {
        if info, err := entry.Info(); err == nil && !info.IsDir() {
            files = append(files, FileInfo{info.Name(), info.Size()})
        }
    }

    sort.Slice(files, func(i, j int) bool {
        return files[i].Size > files[j].Size
    })

    for i := 0; i < 5 && i < len(files); i++ {
        fmt.Printf("%s: %.2f MB\n", files[i].Name, float64(files[i].Size)/(1024*1024))
    }
}