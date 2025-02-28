package main

import (
	"UASM/compiler"
	"UASM/models"
	"log"
	"os"
)

func main() {
	outputFile, inputFile := "output.mem", "misc/teste.uasm"
	tokens, err := compiler.Tokenize(inputFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	// _ = outputFile
	// fmt.Printf("Ok! %d tokens found.\n", len(tokens))
	// for _, tk := range tokens {
	// 	print(tk.String())
	// }

	parser := models.NewParser(tokens)
	compiler.ParseAll(&parser)
	err = os.WriteFile(outputFile, parser.WriteProgram(), 0744)
	if err != nil {
		log.Fatal(err.Error())
	}
}
