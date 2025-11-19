package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	sourceDir := `C:\Users\Public\TP_Automatisation\test_zip`
	os.MkdirAll(sourceDir, os.ModePerm)

	// créer fichiers test
	for i := 0; i < 3; i++ {
		f, _ := os.Create(filepath.Join(sourceDir, fmt.Sprintf("file%d.txt", i)))
		f.WriteString("test\n")
		f.Close()
	}

	zipFile := `C:\Users\Public\TP_Automatisation\backup.zip`
	zf, _ := os.Create(zipFile)
	defer zf.Close()
	zipWriter := zip.NewWriter(zf)
	defer zipWriter.Close()

	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			f, _ := os.Open(path)
			defer f.Close()
			w, _ := zipWriter.Create(info.Name())
			io.Copy(w, f)
		}
		return nil
	})
	fmt.Println("Archive créée :", zipFile)
}
