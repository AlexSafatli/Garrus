package bot

import "github.com/nanobox-io/golang-scribble"

var (
	Db *scribble.Driver
)

func LoadJsonDatabase(dir string) (*scribble.Driver, error) {
	if Db != nil {
		return Db, nil
	}
	Db, err := scribble.New(dir, nil)
	return Db, err
}
