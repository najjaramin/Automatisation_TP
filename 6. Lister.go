package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Utilisateurs connectés :")
	cmd := exec.Command("query", "user")
	cmd.Stdout = exec.Stdout
	cmd.Run()
}
