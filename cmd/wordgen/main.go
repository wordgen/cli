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
	"runtime/debug"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/wordgen/wordgen"
	"github.com/wordgen/wordlists"
)

var version = "dev"

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

type config struct {
	wordCase            string
	wordCount           int
	wordSeparator       string
	selectedWordlist    string
	wordlistPath        string
	printVersion        bool
	printWithoutNewline bool
}

func buildVersion() string {
	if version != "dev" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" || info.Main.Version == "(devel)" {
		return version
	}

	return info.Main.Version
}

func parseFlags(args []string) (*config, error) {
	c := &config{}
	fs := flag.NewFlagSet("wordgen", flag.ContinueOnError)

	fs.StringVar(&c.wordCase, "c", "", "")
	fs.StringVar(&c.wordCase, "case", "", "")
	fs.IntVar(&c.wordCount, "w", 1, "")
	fs.IntVar(&c.wordCount, "words", 1, "")
	fs.StringVar(&c.wordSeparator, "s", " ", "")
	fs.StringVar(&c.wordSeparator, "separator", " ", "")
	fs.StringVar(&c.selectedWordlist, "l", "effLarge", "")
	fs.StringVar(&c.selectedWordlist, "list", "effLarge", "")
	fs.StringVar(&c.wordlistPath, "f", "", "")
	fs.StringVar(&c.wordlistPath, "file", "", "")
	fs.BoolVar(&c.printVersion, "v", false, "")
	fs.BoolVar(&c.printVersion, "version", false, "")
	fs.BoolVar(&c.printWithoutNewline, "n", false, "")
	fs.BoolVar(&c.printWithoutNewline, "no-newline", false, "")
	fs.Usage = func() { fmt.Println(usage) }

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return c, nil
}

func setWordlist(c *config) ([]string, error) {
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

func readWordsFromFile(c *config) ([]string, error) {
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

	fileContent, err := os.ReadFile(absPath) // #nosec G304 -- The --file option intentionally accepts an arbitrary user-selected local path.
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return strings.Split(strings.TrimSpace(string(fileContent)), "\n"), nil
}

func main() {
	c, err := parseFlags(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			return
		}
		os.Exit(2)
	}

	if c.printVersion {
		fmt.Println("wordgen", buildVersion())
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
