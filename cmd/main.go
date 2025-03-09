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
	AssemblerTest()
	InterpreterTest()

	// pp ok
	// np ok
	// pn not ok
	// nn not ok
	//
	//ac := int8(-4)
	//adr := int8(-3)
	//
	//// MUL adr
	//acCache0 := ac // 4
	//siAdr := adr   // 3
	//acCache1 := int8(0)
	//alternate := int8(-1)
	//if siAdr < 0 {
	//	alternate = 1
	//	siAdr = (^siAdr) + 1
	//}
	//for siAdr > 0 {
	//	acCache1 += acCache0
	//	siAdr += alternate
	//}
	//println(acCache1)
}

func AssemblerTest() {
	outputFile, inputFile := "misc/output.mem", "misc/test.uasm"
	tokens, err := compiler.Tokenize(inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	_ = outputFile
	//fmt.Printf("Ok! %d tokens found.\n", len(tokens))
	//for i, tk := range tokens {
	//    print(i, " ", tk.String())
	//}

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
