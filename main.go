package main

import (
	"address2cell/address"
	"address2cell/files"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "process" {
		file := os.Args[2]
		address.Process(files.ReadFile(file))
	} else if command == "help" {
		PrintHelp()
	}
}
