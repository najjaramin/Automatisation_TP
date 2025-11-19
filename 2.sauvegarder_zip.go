package main

import (
    "github.com/mholt/archiver/v3"
    "log"
)

func main() {
    sourceDir := "/etc"
    destZip := "/tmp/backup.zip"

    err := archiver.Archive([]string{sourceDir}, destZip)
    if err != nil {
        log.Fatalf("Erreur lors de la création de l'archive: %v", err)
    } else {
        log.Println("Archive créée avec succès:", destZip)
    }
}