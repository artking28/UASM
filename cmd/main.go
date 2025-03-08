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
	//AssemblerTest()

	// pp ok
	// np ok
	// pn not ok
	// nn not ok

	ac := int8(-4)
	adr := int8(-3)

	// MUL adr
	acCache0 := ac // 4
	siAdr := adr   // 3
	acCache1 := int8(0)
	alternate := int8(-1)
	if siAdr < 0 {
		alternate = 1
		siAdr = (^siAdr) + 1
	}
	for siAdr > 0 {
		acCache1 += acCache0
		siAdr += alternate
	}
	println(acCache1)
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

	parser := models.NewParser(inputFile, tokens)
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
