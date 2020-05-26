package main

import (
	"context"
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/northbright/filehashes"
)

type option struct {
	defaultValue bool
	usage        string
	hashFunc     crypto.Hash
}

var (
	options = map[string]option{
		"md5":    option{true, "compute MD5 checksum", crypto.MD5},
		"sha1":   option{true, "compute SHA1 checksum", crypto.SHA1},
		"sha256": option{false, "compute SHA256 checksum", crypto.SHA256},
		"sha512": option{false, "compute SHA512 checksum", crypto.SHA512},
	}

	optionValues = map[string]*bool{}
	usage        = "Usage:\nhashgo [-md5 | -sha1 | -sha256 | -sha512] [File]...\n"

	concurrency = 1
	bufferSize  = filehashes.DefaultBufferSize
)

// initHashFuncArgs initializes the flags of hash functions.
func initHashFuncArgs() {
	for arg, op := range options {
		optionValues[arg] = flag.Bool(arg, op.defaultValue, op.usage)
	}
}

// getHashFuncs returns the hash functions after the arguments are parsed.
func getHashFuncs() []crypto.Hash {
	var hashFuncs []crypto.Hash

	for arg, op := range options {
		if *(optionValues[arg]) == true {
			hashFuncs = append(hashFuncs, op.hashFunc)
		}
	}

	return hashFuncs
}

func main() {
	// Initialize arguments.
	initHashFuncArgs()
	// Parse flags.
	flag.Parse()

	// Get hash functions in non-flag arguments
	files := flag.Args()
	hashFuncs := getHashFuncs()

	// Check files to be hashed.
	if len(files) == 0 {
		fmt.Printf("Please specify one or more file(s).\n")
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}

	// Check hash functions.
	if len(hashFuncs) == 0 {
		fmt.Printf("Please specify one or more hash function(s).\n")
		fmt.Println(usage)
		flag.PrintDefaults()
		return
	}

	// Create context and cancel function.
	ctx, cancel := context.WithCancel(context.Background())

	// Create a new manager to compute file checksums.
	man, ch := filehashes.NewManager(concurrency, bufferSize)

	for _, file := range files {
		req := filehashes.NewRequest(file, hashFuncs, nil)
		man.Start(ctx, req)
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	totalExited := 0
	for {
		select {
		case <-sigint:
			// os.Interrupt, call cancel func.
			fmt.Printf("os.Interrupt received")
			cancel()
		case <-ctx.Done():
			fmt.Printf("exiting...")
			return
		case m := <-ch:
			switch m.Type {
			case filehashes.STARTED:
				fmt.Printf("%v: starting...\n", m.Req.File)
			case filehashes.ERROR:
				fmt.Printf("%v: error: %v\n", m.Req.File, m.Data.(string))
			case filehashes.DONE:
				fmt.Printf("%v: done\n", m.Req.File)
				checksums := m.Data.(map[crypto.Hash]string)
				for f, checksum := range checksums {
					fmt.Printf("%v: %v\n", f, checksum)
				}
			case filehashes.EXITED:
				totalExited++
				// All goroutine exited
				if totalExited >= len(files) {
					cancel()
				}
			}
		}
	}
}