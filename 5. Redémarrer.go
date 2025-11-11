package main

import (
    "log"
    "os/exec"
)

func main() {
    cmd := exec.Command("sudo", "systemctl", "restart", "nginx")
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Erreur lors du redémarrage du service: %v", err)
    } else {
        log.Println("Service nginx redémarré avec succès.")
    }
}