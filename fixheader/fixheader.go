/*
 * This is free and unencumbered software released into the public domain.
 * For more information, please refer to <http://unlicense.org/>
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(0)
	name := os.Args[0]
	if i := strings.LastIndexAny(name, "/\\"); i > -1 {
		name = name[i+1:]
	}
	log.SetPrefix(name + ": ")
}

func main() {
	p := flag.Bool("p", false, "print")
	no := flag.Bool("no", false, "do not write changes (implies p)")
	flag.Parse()
	name := flag.Arg(0)
	if name == "" {
		log.Fatal("no file specified")
	}
	rom, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	sz := len(rom)
	switch {
	case sz < 512:
		log.Fatal("ROM file must be at least 512 ($200) bytes")
	case sz & 1 != 0:
		log.Fatal("ROM file must be an even size")
	}
	var sum uint16
	for i := 512; i+1 < sz; i += 2 {
		sum += (uint16(rom[i]) << 8) | uint16(rom[i+1])
	}
	if *no == false {
		// apply checksum
		rom[398], rom[399] = byte(sum>>8), byte(sum)
		// apply ROM address range
		rom[416], rom[417], rom[418], rom[419] = 0, 0, 0, 0
		rom[420] = byte((sz-1)>>24)
		rom[421] = byte((sz-1)>>16)
		rom[422] = byte((sz-1)>>8)
		rom[423] = byte(sz-1)
	} else {
		*p = true
	}
	if err = os.WriteFile(name, rom, 0666); err != nil {
		log.Fatal(err)
	}
	if *p == true {
		fmt.Printf("ROM range: $%X-$%X\nchecksum:  $%X\n", 0, sz-1, sum)
	}
}
