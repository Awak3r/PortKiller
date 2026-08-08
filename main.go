package main

import (
    "github.com/Awak3r/PortKiller/parser"
    "github.com/Awak3r/PortKiller/utils"
)
func main() {
    p := utils.Collect()
    parser.ArgParse(p)
}