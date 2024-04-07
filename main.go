package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"

	"github.com/AbdulfatahMohammedSheikh/scriper"
)

func main() {
	// TODO: read args

	args := os.Args

	if 1 == len(args) {
		// TODO: print the mainual or help message
		fmt.Println("no args was passed")
	}

	if 2 == len(args) {

		switch args[1] {

		case "--list":

			c := scriper.New()
			c.GetList("script[type=\"application/json\"]")
			c.Visit("https://github.com/github/gitignore/tree/main")

		default:

			lang := args[1]
			r, size := utf8.DecodeRuneInString(lang)
			if r == utf8.RuneError {
				fmt.Println("cound not deocde entered value")
				os.Exit(1)
			}
			lang = string(unicode.ToUpper(r)) + lang[size:]

			c := scriper.New()
			c.GetIgnorFile("script[type=\"application/json\"]")
			ulr := fmt.Sprintf("https://github.com/github/gitignore/blob/main/%s.gitignore", lang)
			c.Visit(ulr)

		}
	} else if 2 < len(args) {
		// TODO: deal with more args
		// more args
		fmt.Println("command only support one language at a time")
		os.Exit(1)
	}
}
