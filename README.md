# wordgen/cli

A CLI program that prints random words. You can customize the word case,
separator, number of words, and wordlist used.

## Installation

### AUR

There are three AUR packages available:

- [![badge-url]][aur-url]: Download a versioned source archive and build locally.
- [![bin-badge-url]][bin-aur-url]: Download a precompiled binary from the releases page.
- [![git-badge-url]][git-aur-url]: Clone the repository and build the latest commit.

### Non-AUR

#### Go Install

If you have Go installed, install the latest version with:

```sh
go install github.com/wordgen/cli/cmd/wordgen@latest
```

Ensure your Go binary directory is in your `PATH`.

#### Release Binary

Each [release] contains several binaries. Download the binary for your
operating system, then put it in your `PATH`.

The release also includes a `SHA512SUMS` file and its detached signature,
`SHA512SUMS.asc`. The signature is created with the wordgen signing key:

```text
070559FD6C2A19F3
```

#### Build From Source

To build from source, install `git`, `go`, and `make`.

Clone the repository, build the binary, and install it:

```sh
git clone https://github.com/wordgen/cli.git
cd cli
make && make install
```

By default, the binary is installed to `$HOME/.local/bin`.

To uninstall:

```sh
make uninstall
```

## Usage

See `wordgen -h` for a full list of options.

Example output:

```
$ wordgen
vixen

$ wordgen -w 5
hunting clock buffoon trodden deflation

$ wordgen -w 5 -s . -c title
Ditch.Dealer.Flammable.Unearth.Yonder
```

## Wordlists

The wordlists used in this program are from the [wordgen/wordlists] repository.
You can switch which wordlist the words are generated from with the `-l, --list`
option.

You can also use a local wordlist file with the `-f, --file` option. The file
must be a text file with one word per line.

## Contributing

When submitting a pull request, please ensure they are directed to the `dev`
branch of the repository.

Ensure your commit messages and pull request titles follow the
[Conventional Commits] specification.

## License

All files in this repository are licensed under the GNU Affero General Public
License v3.0 or later - see the [LICENSE] file for details.

<!-- links -->
[badge-url]: https://img.shields.io/aur/version/wordgen?label=wordgen&logo=arch-linux&style=plastic
[aur-url]: https://aur.archlinux.org/packages/wordgen
[bin-badge-url]: https://img.shields.io/aur/version/wordgen-bin?label=wordgen-bin&logo=arch-linux&style=plastic
[bin-aur-url]: https://aur.archlinux.org/packages/wordgen-bin
[git-badge-url]: https://img.shields.io/aur/version/wordgen-git?label=wordgen-git&logo=arch-linux&style=plastic
[git-aur-url]: https://aur.archlinux.org/packages/wordgen-git

[release]: https://github.com/wordgen/cli/releases/latest
[wordgen/wordlists]: https://github.com/wordgen/wordlists/blob/main/README.md#available-wordlists
[Conventional Commits]: https://conventionalcommits.org
[LICENSE]: LICENSE
