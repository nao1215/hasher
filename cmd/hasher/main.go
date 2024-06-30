// Package main provides a simple tool to generate hash values.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/nao1215/hasher"
)

// Version value is set by ldflags
var Version string

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

// run is the entry point of this application.
func run() error {
	help := flag.Bool("help", false, "show help message")

	flag.Parse()
	if *help {
		printHelp()
		return nil
	}

	if flag.NArg() < 1 {
		printHelp()
		return errors.New("few arguments")
	}

	h := newHash()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	if err := h.generate(file); err != nil {
		return fmt.Errorf("error generating hash: %w", err)
	}
	return nil
}

// printHelp prints the help message.
func printHelp() {
	fmt.Printf("hasher - A simple tool to generate hash values (version: %s)\n", Version)
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  hasher FILE_PATH")
	fmt.Println("")
	fmt.Println("Limitation:")
	fmt.Println("  hasher command generates hash values from the file.")
}

// hash is a struct that contains the hash map.
type hash struct {
	hash map[string]*hasher.Hash
}

// newHash returns a new hash struct.
func newHash() *hash {
	return &hash{
		hash: map[string]*hasher.Hash{
			"md5sum": hasher.NewHash(hasher.WithMd5()),
			"sha1":   hasher.NewHash(hasher.WithSha1()),
			"sha256": hasher.NewHash(hasher.WithSha256()),
			"sha512": hasher.NewHash(hasher.WithSha512()),
			//"phash":     hasher.NewHash(hasher.WithPhash()),
			"fnv32":   hasher.NewHash(hasher.WithFnv32()),
			"fnv32a":  hasher.NewHash(hasher.WithFnv32a()),
			"fnv64":   hasher.NewHash(hasher.WithFnv64()),
			"fnv64a":  hasher.NewHash(hasher.WithFnv64a()),
			"fnv128":  hasher.NewHash(hasher.WithFnv128()),
			"fnv128a": hasher.NewHash(hasher.WithFnv128a()),
			"blake3":  hasher.NewHash(hasher.WithBlake3()),
			"adler32": hasher.NewHash(hasher.WithAdler32()),
			//"mmh3":      hasher.NewHash(hasher.WithMmh3()),
			//"crc32":     hasher.NewHash(hasher.WithCRC32()),
			"whirlpool": hasher.NewHash(hasher.WithWhirlpool()),
			"xxhash":    hasher.NewHash(hasher.WithXXHash()),
		},
	}
}

// generate generates a hash from the input.
func (h *hash) generate(input io.Reader) error {
	jsonMap := make(map[string]string)

	var buf bytes.Buffer
	_, err := io.Copy(&buf, input)
	if err != nil {
		return err
	}

	for name, hasher := range h.hash {
		hash, err := hasher.Generate(&buf)
		if err != nil {
			return err
		}
		jsonMap[name] = fmt.Sprintf("%x", hash)
	}

	j, err := json.Marshal(jsonMap)
	if err != nil {
		return err
	}

	fmt.Println(string(j))
	return nil
}
