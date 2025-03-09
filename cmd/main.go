package main

import (
	"UASM/compiler"
	"UASM/models"
	"UASM/neander"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	count := len(os.Args)
	if count < 2 || count > 3 {
		log.Fatal(errors.New("error: inválid arguments"))
	}

	inputFile := os.Args[1]
	if !strings.HasSuffix(inputFile, ".uasm") {
		log.Fatal(errors.New("error: this is not a UASM file. Please, rename it before compiling"))
	}

	outputFile := strings.Split(inputFile, ".uasm")[0] + ".mem"
	if count == 3 {
		outputFile = os.Args[2]
		if !strings.HasSuffix(outputFile, ".mem") {
			log.Fatal(errors.New("error: this is not a MEM output file. Please, choose another name to output file"))
		}
	}
	tokens, err := compiler.Tokenize(inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	parser := models.NewParser(inputFile, tokens)
	err = compiler.ParseAll(&parser)
	if err != nil {
		log.Fatal(err.Error())
	}

	content, err := parser.WriteProgram()
	if err != nil {
		log.Fatal(err.Error())
	}

	//parser.Inspect()
	err = os.WriteFile(outputFile, content, 0744)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func InterpreterTest() {
	bytes, err := os.ReadFile("misc/output.mem")
	if err != nil {
		log.Fatalf(err.Error())
	}
	neander.PrintProgram(bytes, false, false, false)

	pr, _ := neander.RunProgram(bytes, false, false)
	fmt.Printf("\n\nResult:\n\tAc = %d, Pc = %d, Z = %v, N = %v\n\n", int8(pr.Ac), pr.Pc, pr.Ac == 0, int8(pr.Ac) < 0)
}
