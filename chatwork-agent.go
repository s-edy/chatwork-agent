package main

import (
	"./chatwork"
)

func main() {
	command := chatwork.NewCommand()
	command.Run()
}
