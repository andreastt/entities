// Converts characters to HTML entities, and vice versa.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andreastt/entities"
)

var reverse = flag.Bool("r", false, "reverses conversion and unescapes input")

func die(err error) {
	s := err.Error()
	fmt.Fprintf(os.Stderr, "%s: error: %s", os.Args[0], s)
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `usage: %s [flags]

Escapes special characters like "<" to become "&lt;".  It currently only
escapes five such characters: <, >, &, ', and ".

flags:
`, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	bs, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		die(err)
	}

	s := string(bs)
	var es string

	if *reverse {
		es = entities.Unescape(s)
	} else {
		es = entities.Escape(s)
	}
	fmt.Printf("%s", es)
}
