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
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/wordgen/wordgen"
	"github.com/wordgen/wordlists"
)

const (
	version = "v0.4.0"
	usage   = `Usage: wordgen [options]

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
)

type Config struct {
	wordCase            string
	wordCount           int
	wordSeparator       string
	selectedWordlist    string
	wordlistPath        string
	printVersion        bool
	printWithoutNewline bool
}

func parseFlags() *Config {
	c := &Config{}

	flag.StringVar(&c.wordCase, "c", "", "")
	flag.StringVar(&c.wordCase, "case", "", "")
	flag.IntVar(&c.wordCount, "w", 1, "")
	flag.IntVar(&c.wordCount, "words", 1, "")
	flag.StringVar(&c.wordSeparator, "s", " ", "")
	flag.StringVar(&c.wordSeparator, "separator", " ", "")
	flag.StringVar(&c.selectedWordlist, "l", "effLarge", "")
	flag.StringVar(&c.selectedWordlist, "list", "effLarge", "")
	flag.StringVar(&c.wordlistPath, "f", "", "")
	flag.StringVar(&c.wordlistPath, "file", "", "")
	flag.BoolVar(&c.printVersion, "v", false, "")
	flag.BoolVar(&c.printVersion, "version", false, "")
	flag.BoolVar(&c.printWithoutNewline, "n", false, "")
	flag.BoolVar(&c.printWithoutNewline, "no-newline", false, "")
	flag.Usage = func() { fmt.Println(usage) }
	flag.Parse()

	return c
}

func setWordlist(c *Config) ([]string, error) {
	if c.wordlistPath != "" {
		return readWordsFromFile(c)
	}

	switch c.selectedWordlist {
	case "effLarge":
		return wordlists.EffLarge, nil
	case "effShort1":
		return wordlists.EffShort1, nil
	case "effShort2":
		return wordlists.EffShort2, nil
	case "namesMixed":
		return wordlists.NamesMixed, nil
	case "namesFemale":
		return wordlists.NamesFemale, nil
	case "namesMale":
		return wordlists.NamesMale, nil
	default:
		return nil, fmt.Errorf("invalid wordlist: %s", c.selectedWordlist)
	}
}

func readWordsFromFile(c *Config) ([]string, error) {
	absPath, err := filepath.Abs(c.wordlistPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	mtype, err := mimetype.DetectFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to detect file type: %v", err)
	}

	if !mtype.Is("text/plain") {
		return nil, fmt.Errorf("not a text file: %s", mtype.String())
	}

	fileContent, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return strings.Split(strings.TrimSpace(string(fileContent)), "\n"), nil
}

func main() {
	c := parseFlags()

	if c.printVersion {
		fmt.Println("wordgen", version)
		return
	}

	wordlist, err := setWordlist(c)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	g := wordgen.NewGenerator()
	g.Words = wordlist
	g.Count = c.wordCount
	g.Casing = c.wordCase
	g.Separator = c.wordSeparator

	words, err := g.Generate()
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	if c.printWithoutNewline {
		fmt.Print(words)
	} else {
		fmt.Println(words)
	}
}
