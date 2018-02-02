package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func main() {
	i := flag.Int("i", 6, "Max Whitespace")
	tabs := flag.Bool("t", false, "Tabbify instead of spaces")
	upper := flag.Bool("u", false, "Uppercase Everything")
	flag.Parse()
	a := flag.Args()
	if flag.NArg() == 0 {
		fmt.Println("provide a string to whitespace")
		os.Exit(1)
	}
	s := arrayJoin(a)
	annoying := randWhiteSpace(s, *i, 0)
	if *upper {
		annoying = strings.ToUpper(annoying)
	}
	if *tabs {
		fmt.Println(tabbify(annoying))
	} else {
		fmt.Println(annoying)
	}
}
