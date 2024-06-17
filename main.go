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
	"fmt"
	"log"

	"github.com/wordgen/wordgen"
)

var version = "dev"

func main() {
	config := parseFlags()

	if config.printVersion {
		fmt.Println("wordgen", version)
		return
	}

	wordlist, err := setWordlist(config)

	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	g := wordgen.NewGenerator()
	g.Words = wordlist
	g.Count = config.wordCount
	g.Casing = config.wordCase
	g.Separator = config.wordSeparator
	words, err := g.Generate()

	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	if config.printWithoutNewline {
		fmt.Print(words)
	} else {
		fmt.Println(words)
	}
}
