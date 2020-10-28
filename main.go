package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}
}

// run is a testable version of main
func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	var (
		i     = flags.Int("i", 6, "Max Whitespace")
		tabs  = flags.Bool("t", false, "Tabbify instead of spaces")
		upper = flags.Bool("u", false, "Uppercase Everything")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}
	a := flags.Args()
	if flags.NArg() == 0 {
		return errors.New("you must provide a string to whitespace")
	}

	s := arrayJoin(a)
	annoying := randWhiteSpace(s, *i, 0)

	if *upper {
		annoying = strings.ToUpper(annoying)
	}
	if *tabs {
		annoying = tabbify(annoying)
	}

	_, _ = fmt.Fprint(stdout, annoying)
	return nil
}

// arrayJoin take an array of strings and return a single one
func arrayJoin(a []string) string {
	var buffer bytes.Buffer
	for _, v := range a {
		buffer.WriteString(v)
	}
	return buffer.String()
}

// randWhiteSpace takes a string, max whitespace between, and an optional random seed and inserts random whitespace
func randWhiteSpace(s string, i int, r int64) string {
	if r == 0 {
		r = time.Now().UnixNano()
	}
	rand.Seed(r)

	var buffer bytes.Buffer
	b := []rune(s)

	for _, v := range b {
		buffer.WriteRune(v)
		c := rand.Intn(i)
		for ii := 0; ii < c; ii++ {
			buffer.WriteString(" ")
		}
	}

	return strings.TrimSpace(buffer.String())
}

// tabbify replaces all spaces in a string with tabs
func tabbify(s string) string {
	r := strings.NewReplacer(" ", "	")
	t := r.Replace(s)
	return t
}
