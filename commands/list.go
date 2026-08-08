package commands

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"PortKiller/utils"
)

func printTable(p []utils.ProcessInfo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintln(w, "PROCESS\tPORT\tPID")
    fmt.Fprintln(w, "-------\t----\t----")
	for _, proc := range p {
		fmt.Fprintf(w, "%s\t%d\t%d\n", proc.Name, proc.Port, proc.Pid)
	}
	w.Flush()
}

func ListAll(p []utils.ProcessInfo) {
	printTable(p)
}

func ListByName(name string, p []utils.ProcessInfo) {
	res := []utils.ProcessInfo{}
	for _, proc := range p {
		if strings.HasPrefix(strings.ToLower(proc.Name),  strings.ToLower(name)) {
			res = append(res, proc)
		}
	}
	printTable(res)
}

func ListByPort(port int, p []utils.ProcessInfo) {
	res := []utils.ProcessInfo{}
	for _, proc := range p {
		if port == proc.Port {
			res = append(res, proc)
		}
	}
	printTable(res)
}