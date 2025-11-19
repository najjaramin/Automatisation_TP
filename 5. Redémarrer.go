package main

import (
	"fmt"
	"os/exec"
)

func main() {
	service := "Spooler"
	fmt.Println("Redémarrage du service", service)
	exec.Command("cmd", "/C", "net stop "+service).Run()
	exec.Command("cmd", "/C", "net start "+service).Run()
	fmt.Println("Service redémarré")
}
