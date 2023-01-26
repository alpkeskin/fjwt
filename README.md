[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
# fjwt
 An another JWT cracker but really fast!!! and simple :)

## Installation

```
go install -v github.com/alpkeskin/fjwt@latest
```

## Usage

```
fjwt eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.mAHLIQfIMpXcoErtJMEeH2eFX93iC3uzbkgrj72OvtY -w WORDLIST_PATH 
```

### Output
```
=====================================
[FOUND] Secret: securepsw1
Attempts: 815176
Elapsed: 2s
=====================================
```

### `--help`

```
An another JWT cracker but really fast!

Usage:
  fjwt [JWT] -w [WORDLIST]

Examples:
fjwt ey... -w wordlist.txt -t 10

Flags:
  -h, --help              help for fjwt
  -t, --threads int       Number of threads (default 10)
  -v, --version           version for fjwt
  -w, --wordlist string   Wordlist path
```

### Roadmap

- [ ] Pause feature will be added.
- [ ] Alphabet feature will be added for cracking process.
