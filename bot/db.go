package bot

import (
	"github.com/nanobox-io/golang-scribble"
	"log"
	"os"
	"path/filepath"
)

var (
	Db          *scribble.Driver
	collections = []string{"entrance", "sound"}
)

func LoadJsonDatabase(dir string) (*scribble.Driver, error) {
	if Db != nil {
		return Db, nil
	}
	Db, err := scribble.New(dir, nil)
	for _, c := range collections {
		// Make sure a json file exists for all collections
		var jsonPath = filepath.Join(dir, c+".json")
		if _, err := os.Stat(jsonPath); err != nil {
			empty, err := os.Create(jsonPath)
			if err != nil {
				log.Fatal(err)
			}
			_ = empty.Close()
		}
	}
	return Db, err
}
