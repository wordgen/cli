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
	"github.com/wordgen/wordlists/eff"
	"github.com/wordgen/wordlists/names"
)

var version = "dev"

const usage = `Usage: wordgen [options]

Options:
  -c, --case STRING         Specify the case of the words: upper, title, lower
  -h, --help                Display this help message and exit
  -l, --list                Specify the wordlist to use
  -n, --no-newline          Print words without a trailing newline
  -s, --separator STRING    Separate words with the specified string
  -v, --version             Print the version and exit
  -w, --words INT           Number of words to print

Wordlists:
  effLarge     namesMixed
  effShort1    namesFemale
  effShort2    namesMale

  effLarge is the default wordlist

Examples:
  wordgen
  wordgen -w 10
  wordgen -w 10 -s . -c title`

func main() {
	var (
		wordCase            string
		wordCount           int
		wordSeparator       string
		selectedWordList    string
		printVersion        bool
		printWithoutNewline bool
		wordlist            []string
	)

	flag.StringVar(&wordCase, "c", "", "")
	flag.StringVar(&wordCase, "case", "", "")
	flag.IntVar(&wordCount, "w", 1, "")
	flag.IntVar(&wordCount, "words", 1, "")
	flag.StringVar(&wordSeparator, "s", " ", "")
	flag.StringVar(&wordSeparator, "separator", " ", "")
	flag.StringVar(&selectedWordList, "l", "effLarge", "")
	flag.StringVar(&selectedWordList, "list", "effLarge", "")
	flag.BoolVar(&printVersion, "v", false, "")
	flag.BoolVar(&printVersion, "version", false, "")
	flag.BoolVar(&printWithoutNewline, "n", false, "")
	flag.BoolVar(&printWithoutNewline, "no-newline", false, "")
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()

	if printVersion {
		fmt.Printf("wordgen %s\n", version)
		return
	}

	switch selectedWordList {
	case "effLarge":
		wordlist = eff.Large
	case "effShort1":
		wordlist = eff.Short1
	case "effShort2":
		wordlist = eff.Short2
	case "namesMixed":
		wordlist = names.Mixed
	case "namesFemale":
		wordlist = names.Female
	case "namesMale":
		wordlist = names.Male
	default:
		log.Fatalln("ERROR: invalid wordlist:", selectedWordList)
	}

	g := wordgen.NewGenerator()
	g.Words = wordlist
	g.Count = wordCount
	g.Casing = wordCase
	g.Separator = wordSeparator
	words, err := g.Generate()

	if err != nil {
		log.Fatalln(err)
	}

	if printWithoutNewline {
		fmt.Print(words)
	} else {
		fmt.Println(words)
	}
}
