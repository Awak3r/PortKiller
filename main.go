package main

import (
    "PortKiller/parser"
    "PortKiller/utils"
)
func main() {
    p := utils.Collect()
    parser.ArgParse(p)
}