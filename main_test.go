package main

import (
	"bytes"
	"testing"
)

func TestArrayJoin(t *testing.T) {
	var tests = []struct {
		name  string
		value []string
		want  string
	}{
		{"happy path", []string{"this", "is", "an", "array", "of", "strings"}, "thisisanarrayofstrings"},
		{"single string", []string{"single"}, "single"},
		{"empty string", []string{""}, ""},
		{"empty array", []string{}, ""},
		{"with newline", []string{"this", "has", "a", "newline\n"}, "thishasanewline\n"},
		{"non-english", []string{"こんにちは", "world"}, "こんにちはworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arrayJoin(tt.value)
			if got != tt.want {
				t.Errorf("arrayJoin(%v), want = %s, got = %s", tt.value, tt.want, got)
			}
		})
	}
}

func TestRandWhiteSpace(t *testing.T) {
	var tests = []struct {
		value string
		want  string
	}{
		{"nospace", "n o   s    p      a       c e"},
		{"endspace ", "e n   d    s      p       a c     e"},
		{" bothspace ", "b   o    t      h       s p     ace"},
		{"こんにちは", "こ ん   に    ち      は"},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got := randWhiteSpace(tt.value, 8, 42)
			if got != tt.want {
				t.Errorf("randWhiteSpace(%s): want = %s, got = %s", tt.value, tt.want, got)
			}
		})
	}
}

func TestTabbify(t *testing.T) {
	var tests = []struct {
		value string
		want  string
	}{
		{"s p  a  c e", "s	p		a		c	e"},
		{"s p  ac e", "s	p		ac	e"},
		{"こ んに   ち  は", "こ	んに			ち		は"},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got := tabbify(tt.value)
			if got != tt.want {
				t.Errorf("tabbify(%s): want = %s, got = %s", tt.value, tt.want, got)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantStdout string
		wantErr    bool
	}{
		{"empty", []string{"prog"}, "", true},
		{"flags no string", []string{"p", "-i", "5"}, "", true},
		{"string", []string{"prog", "s"}, "s", false},
		{"uppsercase", []string{"prog", "-u", "s"}, "S", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stdout := &bytes.Buffer{}
			err := run(tt.args, stdout)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("run() gotStdout = %v, want %v", gotStdout, tt.wantStdout)
			}
		})
	}
}
