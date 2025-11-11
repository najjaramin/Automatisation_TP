package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    utilisateur := "amin" // Remplace par ton nom d'utilisateur
    cmd := exec.Command("ps", "-u", utilisateur)
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Erreur:", err)
        return
    }

    lignes := strings.Split(string(output), "\n")
    nbProcessus := len(lignes) - 2 // En-tête + ligne vide
    fmt.Printf("Nombre de processus pour l'utilisateur %s : %d\n", utilisateur, nbProcessus)
}