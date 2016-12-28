package main

import "testing"

type joinPair struct {
	value    []string
	expected string
}

var jointests = []joinPair{
	{[]string{"this", "is", "an", "array", "of", "strings"}, "thisisanarrayofstrings"},
	{[]string{"single"}, "single"},
	{[]string{""}, ""},
	{[]string{}, ""},
	{[]string{"this", "has", "a", "newline\n"}, "thishasanewline\n"},
	{[]string{"こんにちは", "world"}, "こんにちはworld"},
}

func TestJoin(t *testing.T) {
	for _, pair := range jointests {
		v := arrayJoin(pair.value)
		if v != pair.expected {
			t.Error(
				"For", pair.value,
				"Expected", pair.expected,
				"Got", v,
			)
		}
	}
}

type randPair struct {
	value    string
	expected string
}

var randtests = []randPair{
	{"nospace", "n o   s    p      a       c e"},
	{"endspace ", "e n   d    s      p       a c     e"},
	{" bothspace ", "b   o    t      h       s p     ace"},
	{"こんにちは", "こ ん   に    ち      は"},
}

func TestRand(t *testing.T) {
	for _, pair := range randtests {
		v := randWhiteSpace(pair.value, 8, 42)
		if v != pair.expected {
			t.Error(
				"For", pair.value,
				"Expected", pair.expected,
				"Got", v,
			)
		}
	}
}

var tabtests = []struct {
	value    string
	expected string
}{
	{"s p  a  c e", "s	p		a		c	e"},
	{"s p  ac e", "s	p		ac	e"},
	{"こ んに   ち  は", "こ	んに			ち		は"},
}

func TestTabbify(t *testing.T) {
	for _, pair := range tabtests {
		v := tabbify(pair.value)
		if v != pair.expected {
			t.Error(
				"For", pair.value,
				"Expected", pair.expected,
				"Got", v,
			)
		}
	}
}
