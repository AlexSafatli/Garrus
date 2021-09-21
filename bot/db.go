package bot

import "github.com/nanobox-io/golang-scribble"

func loadJsonDatabase(dir string) (*scribble.Driver, error) {
	return scribble.New(dir, nil)
}
