package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func die(info ...interface{}) {
	fmt.Fprintln(os.Stderr, info...)
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		die(err)
	}
}

func doxy(s, tag string) string {
	n := strings.Index(s, tag)
	if n < 0 {
		return ""
	}
	s = s[n+len(tag):]
	if len(s) == 0 || s[0] != ' ' && s[0] != '\t' {
		return ""
	}
	return strings.TrimSpace(s)
}

type checkedWriter struct {
	w io.Writer
}

func (c checkedWriter) Write(b []byte) (int, error) {
	n, err := c.w.Write(b)
	checkErr(err)
	return n, nil
}

func (c checkedWriter) WriteString(s string) (int, error) {
	n, err := io.WriteString(c.w, s)
	checkErr(err)
	return n, nil
}

type scanner struct {
	*bufio.Scanner
	Name string
	n    int
}

func newScanner(r io.Reader, name string) *scanner {
	return &scanner{
		Scanner: bufio.NewScanner(r),
		Name:    name,
	}
}

func (s *scanner) Scan() bool {
	s.n++
	return s.Scanner.Scan()
}

func (s *scanner) Die(v ...interface{}) {
	v = append(
		[]interface{}{fmt.Sprintf("%s:%d", s.Name, s.n)},
		v...,
	)
	die(v)
}

func upperFirst(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError || unicode.IsUpper(r) {
		return s
	}
	s = s[n:]
	buf := make([]byte, 4+len(s))
	n = utf8.EncodeRune(buf, unicode.ToUpper(r))
	n += copy(buf[n:], s)
	return string(buf[:n])
}