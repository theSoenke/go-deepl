package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thesoenke/go-deepl"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Invalid number of arguments")
	}

	sentence := os.Args[1]
	fromLang := os.Args[2]
	toLang := os.Args[3]
	translations, err := deepl.Translate(sentence, fromLang, toLang)
	if err != nil {
		log.Fatal(err)
	}

	for _, translation := range translations {
		fmt.Printf("%f: %s\n", translation.Probability, translation.Text)
	}
}
