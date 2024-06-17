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
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/wordgen/wordlists/eff"
	"github.com/wordgen/wordlists/names"
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
		return nil, fmt.Errorf("invalid wordlist: %s", c.selectedWordlist)
	}
}

func readWordsFromFile(c Config) ([]string, error) {
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
