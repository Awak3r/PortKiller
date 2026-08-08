package main

import (
	"fmt"
	"github.com/Awak3r/PortKiller/parser"
	"github.com/Awak3r/PortKiller/utils"
	"os"
	"os/exec"
)

func main() {
	if os.Geteuid() != 0 {
		args := os.Args
		cmd := exec.Command("sudo", args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Требуется sudo для работы PortKiller")
			os.Exit(1)
		}
		os.Exit(0)
	}
	p := utils.Collect()
	parser.ArgParse(p)
}
