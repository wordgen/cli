# wordgen/cli

A program that prints random words generated from the [EFF Large Wordlist].

## Installation

### AUR

You can use your preferred AUR helper, or manually clone and run `makepkg`, to
install `wordgen` or `wordgen-bin`.

[![wordgen][badge-url]][aur-url]
[![wordgen-bin][badge-url-bin]][aur-url-bin]

### Non-AUR

#### Release Binary

Each [release] contains several binaries. Download the binary for your
operating system, then put it in your `PATH`.

#### Build From Source

If you wish to build the binary from source, you will need a few things
installed: `git go make`

Clone the repository and run `make` inside the directory:

```
git clone https://github.com/wordgen/cli.git
cd wordgen
make
```

Then run `make install` with root privileges to install:

```
sudo make install
```

To uninstall:

```
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

## Reproducible Builds

Each release binary is built using the [Dockerfile] and [Makefile]. This
enables you to easily verify the release binaries by building one yourself and
comparing the checksums.

For example, to build a Linux amd64 binary, run the following command in the
repository directory:

```
make linux-amd64
```

The binary will be located in the `bin` directory.

## Contributing

When submitting a pull request, please ensure they are directed to the `dev`
branch of the repository.

Ensure your commit messages and pull request titles follow the
[Conventional Commits] specification.

## License

All files in this repository are licensed under the GNU Affero General Public
License v3.0 or later - see the [LICENSE] file for details.

<!-- links -->
[EFF Large Wordlist]: https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt
[badge-url]: https://img.shields.io/aur/version/wordgen?label=wordgen&logo=arch-linux&style=plastic
[aur-url]: https://aur.archlinux.org/packages/wordgen
[badge-url-bin]: https://img.shields.io/aur/version/wordgen-bin?label=wordgen-bin&logo=arch-linux&style=plastic
[aur-url-bin]: https://aur.archlinux.org/packages/wordgen-bin
[release]: https://github.com/wordgen/cli/releases/latest
[Dockerfile]: Dockerfile
[Makefile]: Makefile
[Conventional Commits]: https://conventionalcommits.org
[LICENSE]: LICENSE
