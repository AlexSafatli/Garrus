package bot

import "github.com/nanobox-io/golang-scribble"

func LoadJsonDatabase(dir string) (*scribble.Driver, error) {
	return scribble.New(dir, nil)
}
