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
)

const usage = `Usage: wordgen [options]

Options:
  -c, --case STRING         Specify the case of the words (upper, title, lower)
  -f, --file PATH           Specify a file to use as the wordlist
  -h, --help                Display this help message and exit
  -l, --list STRING         Specify the wordlist to use
  -n, --no-newline          Print words without a trailing newline
  -s, --separator STRING    Separate words with the specified string
  -v, --version             Print the version and exit
  -w, --words INT           Number of words to print

Wordlists:
  effLarge     namesMixed
  effShort1    namesFemale
  effShort2    namesMale

  effLarge is the default wordlist`

type Config struct {
	wordCase            string
	wordCount           int
	wordSeparator       string
	selectedWordlist    string
	wordlistPath        string
	printVersion        bool
	printWithoutNewline bool
}

func parseFlags() Config {
	config := Config{}

	flag.StringVar(&config.wordCase, "c", "", "")
	flag.StringVar(&config.wordCase, "case", "", "")
	flag.IntVar(&config.wordCount, "w", 1, "")
	flag.IntVar(&config.wordCount, "words", 1, "")
	flag.StringVar(&config.wordSeparator, "s", " ", "")
	flag.StringVar(&config.wordSeparator, "separator", " ", "")
	flag.StringVar(&config.selectedWordlist, "l", "effLarge", "")
	flag.StringVar(&config.selectedWordlist, "list", "effLarge", "")
	flag.StringVar(&config.wordlistPath, "f", "", "")
	flag.StringVar(&config.wordlistPath, "file", "", "")
	flag.BoolVar(&config.printVersion, "v", false, "")
	flag.BoolVar(&config.printVersion, "version", false, "")
	flag.BoolVar(&config.printWithoutNewline, "n", false, "")
	flag.BoolVar(&config.printWithoutNewline, "no-newline", false, "")
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()

	return config
}
