package main

import (
	"UASM/neander"
	"fmt"
	"os"
)

func main() {

	bytes, err := os.ReadFile("misc/program.mem")
	if err != nil {
		panic(err.Error())
	}

	result, program := neander.RunProgram(bytes)
	if program == nil {
		panic("something went wrong during the runtime.\n")
	}

	fmt.Printf("success! {ac: %d, pc: %d}\n", result.Ac, result.Pc)
}
