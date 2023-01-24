/*
 * This is free and unencumbered software released into the public domain.
 * For more information, please refer to <http://unlicense.org/>
 */

package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
)

func init() {
	log.SetFlags(0)
	name := os.Args[0]
	if i := strings.IndexAny(name, "/\\"); i > -1 {
		name = name[i+1:]
	}
	log.SetPrefix(name + ": ")
}

func main() {
	// Validate input
	n := flag.Int64("n", 2, "multiple to pad to")
	r := flag.Bool("r", false, "pad with random data")
	flag.Parse()
	name := flag.Arg(0)
	if name == "" {
		log.Fatal("no file specified")
	}
	if *n < 0 {
		log.Fatal("n must be greater than 0")
	}

	// Open target file
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	info, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Generate and write padding
	mod := info.Size() % *n
	if mod == 0 {
		return
	}
	pad := make([]byte, *n-mod)
	if *r == true {
		// fill buffer with random junk data
		for i := 0; i < len(pad); i++ {
			pad[i] = byte(rand.Intn(255))
		}
	}
	f.Write(pad)
}
