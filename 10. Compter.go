package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	user := "ton_nom_utilisateur" // remplacer par os.Getenv("USERNAME")
	out, _ := exec.Command("tasklist", "/FI", "USERNAME eq "+user).Output()
	lines := strings.Split(string(out), "\n")
	count := len(lines) - 3 // enlever en-tête
	fmt.Printf("%d processus pour %s\n", count, user)
}
