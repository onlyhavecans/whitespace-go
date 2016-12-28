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

func arrayJoin(a []string) string {
	var buffer bytes.Buffer
	for _, v := range a {
		buffer.WriteString(v)
	}
	return buffer.String()
}

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

func tabbify(s string) string {
	r := strings.NewReplacer(" ", "	")
	t := r.Replace(s)
	return t
}

func main() {
	i := flag.Int("i", 6, "Max Whitespace")
	tabs := flag.Bool("t", false, "Tabbify instead of spaces")
	flag.Parse()
	a := flag.Args()
	if flag.NArg() == 0 {
		fmt.Println("provide a string to whitespace")
		os.Exit(1)
	}
	s := arrayJoin(a)
	annoying := randWhiteSpace(s, *i, 0)
	if *tabs {
		fmt.Println(tabbify(annoying))
	} else {
		fmt.Println(annoying)
	}
}
