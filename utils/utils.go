package utils

import (
	"log"
	"sort"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	Name string
	Pid int
	Port int
}

func KillByPid(pid int32) bool{
	proc_to_del, err := process.NewProcess(pid)
	if err != nil {
		return false
	}
	err = proc_to_del.Terminate()
	if err == nil {
		return false
	}
	return true
}

func Collect() []ProcessInfo{
	processes := []ProcessInfo{}
	seen := make(map[struct{ pid, port int }]struct{})
	conns, err := net.Connections("tcp4")
    if err != nil {
        log.Fatal(err)
    }
	for _, conn := range conns {
        if conn.Status != "LISTEN" {
            continue 
        }
        p, err := process.NewProcess(conn.Pid)
        if err != nil {
            continue
        }
        name, err := p.Name()
        if err != nil {
            name = "unknown"
        }
		port:= int(conn.Laddr.Port)
		pid := int(conn.Pid)
		k := struct{ pid, port int }{pid, port}
		if _, dup := seen[k]; dup {
			continue
		}
		seen[k] = struct{}{}
		cur:= ProcessInfo{Name: name, Pid: pid, Port: port}
		processes = append(processes,  cur)
    }
	sort.Slice(processes, func(i, j int) bool { return processes[i].Port < processes[j].Port })
	return processes
}