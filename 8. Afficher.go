package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type FileSize struct {
	Name string
	Size int64
}

func main() {
	folder := `C:\Users\Public\TP_Automatisation\test_logs`
	files, _ := ioutil.ReadDir(folder)
	var fs []FileSize
	for _, f := range files {
		if !f.IsDir() {
			fs = append(fs, FileSize{f.Name(), f.Size()})
		}
	}
	sort.Slice(fs, func(i, j int) bool { return fs[i].Size > fs[j].Size })
	for i, f := range fs {
		if i >= 5 {
			break
		}
		fmt.Printf("%s : %d octets\n", f.Name, f.Size)
	}
}
