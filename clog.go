// Copyright 2024 Matthew P. Dargan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Clog creates date-stamped console logs.
//
// Usage:
//
//	clog console logfile
//
// Clog opens the file console and writes every line read from it, prefixed by
// the ASCII time, to the file logfile.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: clog console logfile\n")
	os.Exit(2)
}

func main() {
	log.SetPrefix("clog: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 2 {
		usage()
	}
	con := flag.Arg(0)
	if con == "-" {
		con = os.Stdin.Name()
	}
	cf, err := os.Open(con)
	if err != nil {
		log.Fatal(err)
	}
	defer cf.Close()
	lf, err := os.OpenFile(flag.Arg(1), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		log.Fatal(err)
	}
	defer lf.Close()
	if _, err = lf.Seek(0, io.SeekEnd); err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(cf)
	for s.Scan() {
		l := s.Text()
		logLine := time.Now().Format(time.UnixDate) + ": " + l
		if _, err := fmt.Fprintln(lf, logLine); err != nil {
			log.Fatal(err)
		}
	}
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}
}
