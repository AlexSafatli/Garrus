package sound

import (
	"encoding/json"
	scribble "github.com/nanobox-io/golang-scribble"
	"log"
)

type File struct {
	ID                 string
	FilePath           string
	Categories         []string
	NumberPlays        uint
	ExcludedFromRandom bool
}

type Category struct {
	Name     string
	Children []Category
}

func (l *Library) SetSoundData(sound File, db *scribble.Driver) error {
	if err := db.Write("sound", sound.ID, sound); err != nil {
		return err
	}
	existing := l.SoundMap[sound.ID]
	existing.NumberPlays = sound.NumberPlays
	existing.ExcludedFromRandom = sound.ExcludedFromRandom
	l.SoundMap[sound.ID] = existing
	return nil
}

func (l *Library) LoadSoundData(db *scribble.Driver) error {
	var err error
	var s []string
	s, err = db.ReadAll("sound")
	for _, r := range s {
		loadedSound := File{}
		if err = json.Unmarshal([]byte(r), &loadedSound); err != nil {
			log.Fatalln("Failure loading sounds")
			return err
		}
		if existing, ok := l.SoundMap[loadedSound.ID]; !ok {
			continue
		} else {
			if loadedSound.FilePath != existing.FilePath {
				loadedSound.FilePath = existing.FilePath
				if err := db.Write("sound", loadedSound.ID, loadedSound); err != nil {
					return err
				}
			}
			existing.NumberPlays = loadedSound.NumberPlays
			existing.ExcludedFromRandom = loadedSound.ExcludedFromRandom
			l.SoundMap[loadedSound.ID] = loadedSound
		}
	}
	return err
}
