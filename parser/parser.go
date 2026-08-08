package parser

import (
    "flag"
    "fmt"
    "os"

    "PortKiller/commands"
    "PortKiller/internal/version"
)

func ArgParse() {
    if len(os.Args) < 2 {
        commands.PrintUsage()
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
            commands.ListByName(*name)
        case *port != 0:
            commands.ListByPort(*port)
        default:
            commands.ListAll()
        }

    case "kill":
        cmd := flag.NewFlagSet("kill", flag.ExitOnError)
        name := cmd.String("name", "", "имя процесса")
        port := cmd.Int("port", 0, "порт")
        cmd.Parse(os.Args[2:])

        switch {
        case *name != "":
            commands.KillByName(*name)
        case *port != 0:
            commands.KillByPort(*port)
        default:
            fmt.Println("укажи -name или -port")
        }

    default:
        fmt.Println("неизвестная команда:", os.Args[1])
        commands.PrintUsage()
    }
}