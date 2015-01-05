package do_test

import (
	"testing"

	"github.com/cheekybits/godo/do"
	"github.com/cheekybits/is"
)

func TestDo(t *testing.T) {
	is := is.New(t)

	d := do.New()
	d.Tokens = []string{"TODO"}
	var locations []*do.Location
	for r := range d.Walk("./test", "*.go") {
		locations = append(locations, r)
	}

	is.Equal(len(locations), 2)

}

func TestLocationString(t *testing.T) {
	is := is.New(t)

	l := &do.Location{File: "file.go", Line: 2, Preview: "// TODO: remember to do something", Token: "TODO"}
	is.Equal(l.String(), "file.go:2: // TODO: remember to do something")

}
