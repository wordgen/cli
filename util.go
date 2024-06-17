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
	"os"
	"strings"

	"github.com/wordgen/wordlists/eff"
	"github.com/wordgen/wordlists/names"
	"golang.org/x/tools/godoc/util"
	"golang.org/x/tools/godoc/vfs"
)

func setWordlist(c Config) ([]string, error) {
	if c.wordlistPath != "" {
		return readWordsFromFile(c)
	}

	switch c.selectedWordlist {
	case "effLarge":
		return eff.Large, nil
	case "effShort1":
		return eff.Short1, nil
	case "effShort2":
		return eff.Short2, nil
	case "namesMixed":
		return names.Mixed, nil
	case "namesFemale":
		return names.Female, nil
	case "namesMale":
		return names.Male, nil
	default:
		return []string{}, fmt.Errorf("invalid wordlist: %s", c.selectedWordlist)
	}
}

func readWordsFromFile(c Config) ([]string, error) {
	if !util.IsTextFile(vfs.OS("."), c.wordlistPath) {
		return []string{}, fmt.Errorf("invalid file path: %s", c.wordlistPath)
	}

	fileContent, err := os.ReadFile(c.wordlistPath)

	if err != nil {
		return []string{}, fmt.Errorf("failed to read file: %v", err)
	}

	return strings.Split(strings.TrimSpace(string(fileContent)), "\n"), nil
}
