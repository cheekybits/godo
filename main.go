package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/cheekybits/godo/do"
)

/*

	Usage

		godo -tokens=TODO {path}

*/

func main() {
	var (
		tokens  = flag.String("tokens", "TODO,FIXME", "list of comma separated tokens")
		pattern = flag.String("pattern", "*.*", "Pattern of files to process")
	)
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"./"}
	}
	d := do.New()
	d.Tokens = strings.Split(*tokens, ",")
	for item := range d.Walk(args[0], *pattern) {
		fmt.Println(item.String())
	}
	if d.Err != nil {
		log.Fatalln(d.Err)
	}
}
