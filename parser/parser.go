package parser

import (
    "flag"
    "fmt"
    "os"
    "github.com/Awak3r/PortKiller/utils"
    "github.com/Awak3r/PortKiller/commands"
    "github.com/Awak3r/PortKiller/internal/version"
)

func ArgParse(proc []utils.ProcessInfo) {
    if len(os.Args) < 2 {
        fmt.Println("run with --help for usage")
        return
    }
    switch os.Args[1] {
    case "--version", "-v", "version":
        fmt.Println(version.Full())

    case "--help", "-help", "-h":
        commands.PrintUsage()

    case "list":
        cmd := flag.NewFlagSet("list", flag.ExitOnError)
        name := cmd.String("name", "", "имя процесса")
        port := cmd.Int("port", 0, "порт")
        cmd.Parse(os.Args[2:])

        switch {
        case *name != "":
            commands.ListByName(*name, proc)
        case *port != 0:
            commands.ListByPort(*port, proc)
        default:
            commands.ListAll(proc)
        }

    case "kill":
        cmd := flag.NewFlagSet("kill", flag.ExitOnError)
        name := cmd.String("name", "", "имя процесса")
        port := cmd.Int("port", 0, "порт")
        cmd.Parse(os.Args[2:])

        switch {
        case *name != "":
            res := commands.KillByName(*name, proc)
            if (res == false) {
                fmt.Println("Не удалось завершить процесс")
            }
        case *port != 0:
            res := commands.KillByPort(*port, proc)
            if (res == false) {
                fmt.Println("Не удалось завершить процесс")
            }
        default:
            fmt.Println("укажи -name или -port")
        }

    default:
        fmt.Println("неизвестная команда:", os.Args[1])
        commands.PrintUsage()
    }
}