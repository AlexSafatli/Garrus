package sound

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
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

func (l *Library) SetSoundData(sound *File, db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sound"))
		if b == nil {
			return fmt.Errorf("get bucket: %+v", b)
		}
		buf, err := json.Marshal(*sound)
		if err != nil {
			return err
		}
		return b.Put([]byte(sound.ID), buf)
	})
	if err != nil {
		return err
	}
	existing := l.SoundMap[sound.ID]
	existing.NumberPlays = sound.NumberPlays
	existing.ExcludedFromRandom = sound.ExcludedFromRandom
	l.SoundMap[sound.ID] = existing
	return nil
}

func (l *Library) LoadSoundData(db *bolt.DB) error {
	var err error
	var s []File
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("sound"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = b.ForEach(func(k, v []byte) error {
			sf := File{}
			if err = json.Unmarshal(v, &sf); err != nil {
				return err
			}
			s = append(s, sf)
			return nil
		})
		return err
	})
	for _, sf := range s {
		if existing, ok := l.SoundMap[sf.ID]; !ok {
			continue
		} else {
			if sf.FilePath != existing.FilePath {
				sf.FilePath = existing.FilePath
				if err := l.SetSoundData(&sf, db); err != nil {
					return err
				}
			}
			existing.NumberPlays = sf.NumberPlays
			existing.ExcludedFromRandom = sf.ExcludedFromRandom
			l.SoundMap[sf.ID] = &sf
		}
	}
	return err
}

func (f *File) ContainsCategory(cat string) bool {
	for _, c := range f.Categories {
		if c == cat {
			return true
		}
	}
	return false
}
