package commands

import ("fmt")

func PrintUsage() {
    fmt.Println(`PortKiller — kill процессов по имени/порту

Usage:
  portkiller list [-name NAME | -port PORT]
  portkiller kill [-name NAME | -port PORT]
  portkiller -version
  portkiller -help`)
}