## Hasher
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
[![Multi OS Unit Test](https://github.com/nao1215/hasher/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/hasher/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/nao1215/hasher/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/hasher/actions/workflows/reviewdog.yml)

The `hasher` package operates on different hash algorithms through a unified interface. The interfaces it provides include "hash generation" and "comparison of hash values with files (or strings)."

The `hasher` is cross-platform and has been tested on Windows, macOS, and Linux.

### Supported Hash Algorithms

- MD5
- SHA1
- SHA256
- SHA512
- 32-bit FNV-1, FNV-1a
- 64-bit FNV-1, FNV-1a
- 128-bit FNV-1, FNV-1a
- Perceptual Hash (only for images) 
- User-defined algorithms

## Usage

### Use default algorithm: MD5

```go
package main

import (
	"fmt"
	"github.com/nao1215/hasher"
	"log"
	"os"
)

func main() {
	// Create a new Hash instance with default options.
	h := hasher.NewHash()

	// Generate a hash from a string.
	hash, err := h.Generate("example")
	if err != nil {
	    log.Fatal(err)
	}

	// Compare a hash with a string.
	err := h.Compare(hash, "example")
	if err != nil {
	    log.Fatal(err)
	}

	// Generate a hash from an io.Reader.
	file, err := os.Open("example.txt")
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()
	hash, err = h.Generate(file)
	if err != nil {
	    log.Fatal(err)
	}

	// Compare a hash with an io.Reader.
	file, err = os.Open("example.txt")
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	err = h.Compare(hash, file)
	if err != nil {
	    log.Fatal(err)
	}
}
```

### Use another algorithm: SHA256

If you use another algorithm, you can specify algorithm option when creating a new Hash instance. If you use SHA256, you can do as follows:

```go
    h := hasher.NewHash(hasher.WithSha256sum())
```

### Use user-defined algorithm

If you use a user-defined algorithm, you must implement the `Hasher` interface.

```go
// Hasher is an interface that contains the methods to generate and compare hashes.
type Hasher interface {
	// GenHashFromString generates a hash from a string.
	GenHashFromString(string) ([]byte, error)
	// GenHashFromIOReader generates a hash from an io.Reader.
	GenHashFromIOReader(io.Reader) ([]byte, error)
	// CmpHashAndString compares a hash and a string.
	// If the hash and the string are the same, nil is returned.
	CmpHashAndString([]byte, string) error
	// CmpHashAndIOReader compares a hash and an io.Reader.
	// If the hash and the io.Reader are the same, nil is returned.
	CmpHashAndIOReader([]byte, io.Reader) error
}
```

```go
	// YourOriginalHashAlgorithm implements the Hasher interface.
	h := hasher.NewHash(hasher.WithUserDifinedAlgorithm(YourOriginalHashAlgorithm))
```

## LICENSE
[MIT License](./LICENSE)

## Contribution
First off, thanks for taking the time to contribute! Contributions are not only related to development. For example, GitHub Star motivates me to develop! Please feel free to contribute to this project.


Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://debimate.jp/"><img src="https://avatars.githubusercontent.com/u/22737008?v=4?s=75" width="75px;" alt="CHIKAMATSU Naohiro"/><br /><sub><b>CHIKAMATSU Naohiro</b></sub></a><br /><a href="https://github.com/nao1215/hasher/commits?author=nao1215" title="Code">💻</a></td>
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td align="center" size="13px" colspan="7">
        <img src="https://raw.githubusercontent.com/all-contributors/all-contributors-cli/1b8533af435da9854653492b1327a23a4dbd0a10/assets/logo-small.svg">
          <a href="https://all-contributors.js.org/docs/en/bot/usage">Add your contributions</a>
        </img>
      </td>
    </tr>
  </tfoot>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
