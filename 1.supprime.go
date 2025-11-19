package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	logDir := `C:\Users\Public\TP_Automatisation\test_logs`
	days := 7
	limit := time.Now().AddDate(0, 0, -days)

	os.MkdirAll(logDir, os.ModePerm)

	filepath.Walk(logDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.ModTime().Before(limit) {
			fmt.Println("Suppression de :", path)
			os.Remove(path)
		}
		return nil
	})
}
