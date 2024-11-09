package main

import (
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"

	"github.com/AbdulfatahMohammedSheikh/scriper"
)

func main() {


	args := os.Args

	if len(args) < 2 {
		fmt.Println("no args was passed use --help to get more inforamtion")
	}

	if 2 == len(args) {

		switch args[1] {

		case "--list":

			c := scriper.New()
			c.GetList("script[type=\"application/json\"]")
			c.OnError()
			c.Visit("https://github.com/github/gitignore/tree/main")

		case "--help":
			fmt.Println("use --list to check if the language is supported")
			fmt.Println("use <lang_name> to get the .gitignore file data")
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
			c.OnError()
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
