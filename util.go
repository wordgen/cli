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
)

func getWordlist(wordlist string) ([]string, error) {
	if isFile(wordlist) {
		return readWordsFromFile(wordlist)
	}

	switch wordlist {
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
		return []string{}, fmt.Errorf("invalid wordlist: %s", wordlist)
	}
}

func readWordsFromFile(path string) ([]string, error) {
	fileContent, err := os.ReadFile(path)

	if err != nil {
		return []string{}, fmt.Errorf("failed to read file: %v", err)
	}

	return strings.Split(strings.TrimSpace(string(fileContent)), "\n"), nil
}

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && fileInfo.Mode().IsRegular()
}
