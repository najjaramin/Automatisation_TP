package main

import (
    "os"
    "path/filepath"
    "time"
)

func main() {
    logDir := "/var/log/myapp"
    cutoff := time.Now().AddDate(0, 0, -7)

    filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
        if err == nil && !info.IsDir() && info.ModTime().Before(cutoff) {
            os.Remove(path)
        }
        return nil
    })
}