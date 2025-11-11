package main

import (
    "os/exec"
)

func main() {
    cmd := exec.Command("bash", "-c", `echo "0 2 * * * /path/to/script.sh" | crontab -`)
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
}