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
	"github.com/wordgen/wordlists"
)

var version = "dev"

func main() {
	var (
		wordCase            = flag.String("c", "lower", "")
		wordCount           = flag.Int("w", 1, "")
		wordSeparator       = flag.String("s", " ", "")
		printVersion        = flag.Bool("v", false, "")
		printWithoutNewline = flag.Bool("n", false, "")
	)

	flag.Usage = func() {
		s := `USAGE
  wordgen
  wordgen [OPTIONS]

OPTIONS
  -c <STRING>  Specify the case of the words: upper, title, lower
               (default lower)

  -h           Display this help message and exit

  -n           Print words without a trailing newline

  -s <STRING>  Separate words with the specified string
               (default " ")

  -v           Print the version and exit

  -w <INT>     Number of words to print
               (default 1)

EXAMPLES
  wordgen
  wordgen -w 10
  wordgen -w 10 -s . -c title`

		fmt.Println(s)
	}

	flag.Parse()

	if *printVersion {
		fmt.Printf("wordgen %s\n", version)
		return
	}

	words, err := wordgen.WordGenerator(wordlists.EFFLarge, *wordCount, *wordCase, *wordSeparator)

	if err != nil {
		log.Fatalln(err)
	}

	if *printWithoutNewline {
		fmt.Print(words)
	} else {
		fmt.Println(words)
	}
}
