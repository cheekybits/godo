package do

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Finder struct {
	Tokens []string
	Walker func(root string, walkFn filepath.WalkFunc) error
	Err    error
}

type Location struct {
	File    string
	Line    int
	Preview string
	Token   string
}

func (l *Location) String() string {
	return fmt.Sprintf("%s:%d: %s", l.File, l.Line, l.Preview)
}

func (f *Finder) Walk(path, pattern string) <-chan *Location {
	tokens := make([]string, len(f.Tokens))
	for i, t := range f.Tokens {
		tokens[i] = strings.ToLower(t)
	}
	c := make(chan *Location)
	go func() {
		f.Err = f.Walker(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if match, err := filepath.Match(pattern, filepath.Base(path)); !match {
				return err
			}
			file, err := os.Open(path)
			s := bufio.NewScanner(file)
			l := 0
			for s.Scan() {
				l++
				for _, tok := range tokens {
					if strings.Contains(strings.ToLower(s.Text()), tok) {
						c <- &Location{File: path, Line: l, Preview: preview(s.Text()), Token: tok}
					}
				}
			}
			file.Close()
			return nil
		})
		close(c)
	}()
	return c
}

func preview(s string) string {
	return strings.Trim(s, "\n \t\r")
}

func New() *Finder {
	return &Finder{
		Tokens: []string{"TODO"},
		Walker: filepath.Walk,
	}
}
