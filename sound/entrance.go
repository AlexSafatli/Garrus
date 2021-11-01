package sound

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type Entrance struct {
	UserID              string
	SoundID             string
	PersonalizedMessage string
}

var entrances []Entrance

func GetEntranceForUser(userID string) *Entrance {
	for _, e := range entrances {
		if e.UserID == userID {
			return &e
		}
	}
	return nil
}

func getEntranceIndex(userID string) int {
	for i, e := range entrances {
		if e.UserID == userID {
			return i
		}
	}
	return -1
}

func SetEntranceForUser(userID, soundID, msg string, db *bolt.DB) error {
	var err error
	e := Entrance{UserID: userID, SoundID: soundID, PersonalizedMessage: msg}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entrance"))
		if b == nil {
			return fmt.Errorf("get bucket: %+v", b)
		}
		buf, err := json.Marshal(e)
		if err != nil {
			return err
		}
		return b.Put([]byte(userID), buf)
	})
	if err != nil {
		return err
	}
	i := getEntranceIndex(userID)
	if i >= 0 {
		log.Printf("Found entrance %+v; changing sound ID", entrances[i])
		entrances[i].SoundID = soundID
	} else {
		entrances = append(entrances, e)
	}
	return nil
}

func DeleteEntranceForUser(userID string, db *bolt.DB) error {
	var err error
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entrance"))
		if b == nil {
			return fmt.Errorf("get bucket: %s", err)
		}
		return b.Delete([]byte(userID))
	})
	if err != nil {
		return err
	}
	e := getEntranceIndex(userID)
	if e != -1 {
		entrances[e] = entrances[len(entrances)-1]
		entrances = entrances[:len(entrances)-1]
	}
	return nil
}

func LoadEntrances(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("entrance"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = b.ForEach(func(k, v []byte) error {
			ent := Entrance{}
			if err = json.Unmarshal(v, &ent); err != nil {
				return err
			}
			entrances = append(entrances, ent)
			return nil
		})
		return err
	})
}
