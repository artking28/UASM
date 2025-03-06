package main

import (
	"UASM/compiler"
	"UASM/models"
	"UASM/neander"
	"fmt"
	"log"
	"os"
)

func main() {
	//InterpreterTest()
	AssemblerTest()
}

func AssemblerTest() {
	outputFile, inputFile := "output.mem", "misc/teste.uasm"
	tokens, err := compiler.Tokenize(inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	_ = outputFile
	//fmt.Printf("Ok! %d tokens found.\n", len(tokens))
	//for i, tk := range tokens {
	//    print(i, " ", tk.String())
	//}

	parser := models.NewParser(tokens)
	err = compiler.ParseAll(&parser)
	if err != nil {
		log.Fatal(err.Error())
	}

	parser.Inspect()
	//err = os.WriteFile(outputFile, parser.WriteProgram(), 0744)
	//if err != nil {
	//    log.Fatal(err.Error())
	//}
}

func InterpreterTest() {
	bytes, err := os.ReadFile("misc/TESTEHEREDIA.mem")
	if err != nil {
		log.Fatal(err.Error())
	}
	//neander.PrintProgram(bytes, false, true)

	pr, _ := neander.RunProgram(bytes, false, true)
	fmt.Printf("\nResult:\n\tAc = %x, Pc = %x, Z = %v, N = %v\n", pr.Ac, pr.Pc, pr.Ac == 0, int8(pr.Ac) < 0)
}
