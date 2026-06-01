package main

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/wordgen/wordlists"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want *config
	}{
		{
			name: "defaults",
			want: &config{
				wordCase:            "",
				wordCount:           1,
				wordSeparator:       " ",
				selectedWordlist:    "effLarge",
				wordlistPath:        "",
				printVersion:        false,
				printWithoutNewline: false,
			},
		},
		{
			name: "short options",
			args: []string{
				"-c", "upper",
				"-w", "3",
				"-s", ":",
				"-l", "namesMale",
				"-f", "words.txt",
				"-v",
				"-n",
			},
			want: &config{
				wordCase:            "upper",
				wordCount:           3,
				wordSeparator:       ":",
				selectedWordlist:    "namesMale",
				wordlistPath:        "words.txt",
				printVersion:        true,
				printWithoutNewline: true,
			},
		},
		{
			name: "long options",
			args: []string{
				"--case", "lower",
				"--words", "5",
				"--separator", ",",
				"--list", "effShort1",
				"--file", "custom.txt",
				"--version",
				"--no-newline",
			},
			want: &config{
				wordCase:            "lower",
				wordCount:           5,
				wordSeparator:       ",",
				selectedWordlist:    "effShort1",
				wordlistPath:        "custom.txt",
				printVersion:        true,
				printWithoutNewline: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFlags(tt.args)
			if err != nil {
				t.Fatalf("parseFlags() error = %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("parseFlags() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestSetWordlist(t *testing.T) {
	tests := []struct {
		name     string
		selected string
		want     []string
		wantErr  bool
	}{
		{
			name:     "eff large",
			selected: "effLarge",
			want:     wordlists.EffLarge,
		},
		{
			name:     "eff short 1",
			selected: "effShort1",
			want:     wordlists.EffShort1,
		},
		{
			name:     "eff short 2",
			selected: "effShort2",
			want:     wordlists.EffShort2,
		},
		{
			name:     "female names",
			selected: "namesFemale",
			want:     wordlists.NamesFemale,
		},
		{
			name:     "male names",
			selected: "namesMale",
			want:     wordlists.NamesMale,
		},
		{
			name:     "mixed names",
			selected: "namesMixed",
			want:     wordlists.NamesMixed,
		},
		{
			name:     "invalid list",
			selected: "invalid",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setWordlist(&config{
				selectedWordlist: tt.selected,
			})

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadWordsFromFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "words.txt")

	if err := os.WriteFile(path, []byte("alpha\nbeta\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := readWordsFromFile(&config{
		wordlistPath: path,
	})
	if err != nil {
		t.Fatal(err)
	}

	want := []string{"alpha", "beta"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
