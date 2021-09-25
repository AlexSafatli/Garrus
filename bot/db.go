package bot

import (
	"github.com/boltdb/bolt"
)

var (
	dbPath string
)

func SetDatabasePath(path string) {
	dbPath = path
}

func LoadDatabase() (*bolt.DB, error) {
	return bolt.Open(dbPath, 0600, nil)
}
