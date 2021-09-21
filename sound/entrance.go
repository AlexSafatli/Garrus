package sound

import (
	"encoding/json"
	scribble "github.com/nanobox-io/golang-scribble"
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

func SetEntranceForUser(userID, soundID, msg string, db *scribble.Driver) error {
	e := Entrance{UserID: userID, SoundID: soundID, PersonalizedMessage: msg}
	if err := db.Write("entrance", userID, e); err != nil {
		return err
	}
	e_ := GetEntranceForUser(userID)
	if e_ != nil {
		e_.SoundID = soundID
	} else {
		entrances = append(entrances, e)
	}
	return nil
}

func DeleteEntranceForUser(userID string, db *scribble.Driver) error {
	err := db.Delete("entrance", userID)
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

func LoadEntrances(db *scribble.Driver) error {
	var err error
	var e []string
	e, err = db.ReadAll("entrance")
	for _, r := range e {
		loadedEntrance := Entrance{}
		if err = json.Unmarshal([]byte(r), &loadedEntrance); err != nil {
			return err
		}
		entrances = append(entrances, loadedEntrance)
	}
	return err
}
