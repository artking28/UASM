package main

import (
	"UASM/compiler"
	"fmt"
	"log"
)

func main() {

	tokens, err := compiler.Tokenize("misc/teste.uasm")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Ok! %d tokens found.\n", len(tokens))
}
