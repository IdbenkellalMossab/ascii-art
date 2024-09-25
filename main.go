package main

import (
	"fmt"
	"os"

	files "ascii_art/ASCII"
)

func main() {
	Standard := files.NewFile("Symbols/standard.txt")
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("The number of arguments is not correct!!!")
		return
	} else {
		if args[0] == "" {
			return
		}
		Standard.SplitAndExec(args[0])
		fmt.Print(Standard.Result)
	}
}
