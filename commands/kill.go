package commands

import (
	"strings"
	"PortKiller/utils"
)

func KillByName(name string, p []utils.ProcessInfo) bool {
    killed := false
    for _, proc := range p {
        if strings.HasPrefix(strings.ToLower(proc.Name), strings.ToLower(name)) {
            killed = utils.KillByPid(int32(proc.Pid))
        }
    }
    return killed
}

func KillByPort(port int, p []utils.ProcessInfo) bool {
    killed := false
    for _, proc := range p {
        if port == proc.Port {
            killed = utils.KillByPid(int32(proc.Pid))
        }
    }
    return killed
}