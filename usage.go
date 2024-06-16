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

const usage = `Usage: wordgen [options]

Options:
  -c, --case STRING         Specify the case of the words (upper, title, lower)
  -h, --help                Display this help message and exit
  -l, --list STRING         Specify the wordlist to use (can be a file)
  -n, --no-newline          Print words without a trailing newline
  -s, --separator STRING    Separate words with the specified string
  -v, --version             Print the version and exit
  -w, --words INT           Number of words to print

Wordlists:
  effLarge     namesMixed
  effShort1    namesFemale
  effShort2    namesMale

  effLarge is the default wordlist`
