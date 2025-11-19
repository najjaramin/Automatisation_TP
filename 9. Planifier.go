package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := `schtasks /Create /SC DAILY /TN TestScript /TR "python C:\Users\Public\TP_Automatisation\1.supprime.py" /ST 02:00 /F`
	err := exec.Command("cmd", "/C", cmd).Run()
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Tâche planifiée à 2h du matin")
	}
}
