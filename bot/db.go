package bot

import (
	"github.com/boltdb/bolt"
	"log"
)

var (
	dbPath string
)

func SetDatabasePath(path string) {
	dbPath = path
}

func LoadDatabase() *bolt.DB {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		log.Fatalln("Could not load database", err)
	}
	return db
}
