// Copyright (C) 2024  Daniel Kuehn <daniel@kuehn.foo>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wordgen/wordgen"
)

var version = "dev"

func main() {
	var (
		wordCase            string
		wordCount           int
		wordSeparator       string
		selectedWordlist    string
		printVersion        bool
		printWithoutNewline bool
	)

	flag.StringVar(&wordCase, "c", "", "")
	flag.StringVar(&wordCase, "case", "", "")
	flag.IntVar(&wordCount, "w", 1, "")
	flag.IntVar(&wordCount, "words", 1, "")
	flag.StringVar(&wordSeparator, "s", " ", "")
	flag.StringVar(&wordSeparator, "separator", " ", "")
	flag.StringVar(&selectedWordlist, "l", "effLarge", "")
	flag.StringVar(&selectedWordlist, "list", "effLarge", "")
	flag.BoolVar(&printVersion, "v", false, "")
	flag.BoolVar(&printVersion, "version", false, "")
	flag.BoolVar(&printWithoutNewline, "n", false, "")
	flag.BoolVar(&printWithoutNewline, "no-newline", false, "")
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()

	if printVersion {
		fmt.Println("wordgen", version)
		return
	}

	wordlist, err := getWordlist(selectedWordlist)

	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	g := wordgen.NewGenerator()
	g.Words = wordlist
	g.Count = wordCount
	g.Casing = wordCase
	g.Separator = wordSeparator
	words, err := g.Generate()

	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	if printWithoutNewline {
		fmt.Print(words)
	} else {
		fmt.Println(words)
	}
}
