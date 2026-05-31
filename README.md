# wordgen/cli

A CLI program that prints random words. You can customize the word case,
separator, number of words, and wordlist used.

## Installation

### AUR

You can use your preferred AUR helper, or manually clone and run `makepkg`, to
install `wordgen` or `wordgen-bin`.

[![wordgen][badge-url]][aur-url]
[![wordgen-bin][badge-url-bin]][aur-url-bin]

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

If you wish to build the binary from source, you will need a few things
installed: `git go make`

Clone the repository and run `make` inside the directory:

```sh
git clone https://github.com/wordgen/cli.git
cd cli
make
```

Then run `make install` with root privileges to install:

```sh
sudo make install
```

To uninstall:

```sh
sudo make uninstall
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
[wordgen/wordlists]: https://github.com/wordgen/wordlists/blob/main/README.md#available-wordlists
[badge-url]: https://img.shields.io/aur/version/wordgen?label=wordgen&logo=arch-linux&style=plastic
[aur-url]: https://aur.archlinux.org/packages/wordgen
[badge-url-bin]: https://img.shields.io/aur/version/wordgen-bin?label=wordgen-bin&logo=arch-linux&style=plastic
[aur-url-bin]: https://aur.archlinux.org/packages/wordgen-bin
[release]: https://github.com/wordgen/cli/releases/latest
[Conventional Commits]: https://conventionalcommits.org
[LICENSE]: LICENSE
