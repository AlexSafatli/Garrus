package sound

import (
	"github.com/AlexSafatli/Garrus/trie"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Library struct {
	RootPath   string
	SoundMap   map[string]*File
	Categories []string
	Trie       *trie.LowercaseTrie
}

var library Library

func (l *Library) doConversions() []error {
	var errs []error
	for _, v := range l.SoundMap {
		if needsConversionToDCA(v.FilePath) {
			var newPath = strings.TrimSuffix(v.FilePath, filepath.Ext(v.FilePath)) + ".dca"
			log.Printf("Converting sound %s to DCA -> %s", v.ID, newPath)
			err := saveSoundFileToDCA(v.FilePath, newPath)
			if err != nil {
				errs = append(errs, err)
			} else {
				err = os.Remove(v.FilePath) // remove the old path
				if err != nil {
					errs = append(errs, err)
				}
			}
			v.FilePath = newPath
		}
	}
	return errs
}

func (l *Library) Category(s string) (bool, string) {
	s = strings.ToLower(s)
	for _, c := range library.Categories {
		if s == strings.ToLower(c) {
			return true, c
		}
	}
	return false, ""
}

func (l *Library) GetSoundNames() (keys []string) {
	keys = make([]string, 0, len(l.SoundMap))
	for k := range l.SoundMap {
		keys = append(keys, k)
	}
	return
}

func (l *Library) Contains(s string) bool {
	_, ok := l.SoundMap[s]
	return ok
}

func (l *Library) GetClosestMatchingSoundID(s string) string {
	return l.Trie.GetWordWithPrefix(s)
}

func (l *Library) GetRandomSound() *File {
	var i uint
	r := uint(rand.Intn(len(l.SoundMap)))
	for _, v := range l.SoundMap {
		if i == r {
			return v
		}
		i++
	}
	return nil
}

func (l *Library) GetRandomSoundForCategory(category string) *File {
	var i uint
	r := uint(rand.Intn(len(l.SoundMap)))
	for i <= uint(len(l.SoundMap))+r {
		for _, v := range l.SoundMap {
			if i >= r && v.ContainsCategory(category) {
				return v
			}
			i++
		}
	}
	return nil
}

func LoadSounds(rootPath string) error {
	l := Library{RootPath: rootPath, SoundMap: make(map[string]*File)}
	files, cats, err := walkRootDirectoryForSounds(rootPath, "")
	if err != nil {
		return err
	}
	for i := range files {
		l.SoundMap[files[i].ID] = &files[i]
	}
	errs := l.doConversions()
	if len(errs) > 0 {
		return errs[0]
	}
	l.Categories = cats
	l.Trie = trie.NewLowercaseTrie(l.GetSoundNames())
	library = l
	return nil
}

func GetLibrary() *Library {
	return &library
}

func walkRootDirectoryForSounds(start, top string) (files []File, categories []string, err error) {
	var cleanStart = filepath.Clean(start)
	var catMap = make(map[string]bool)
	if len(top) == 0 {
		top = cleanStart
	}
	err = filepath.Walk(cleanStart, func(path string, info os.FileInfo, err error) error {
		if err != nil || cleanStart == path || strings.HasPrefix(info.Name(), ".") {
			return nil // ignore hidden files, etc.
		}
		if !info.IsDir() && isSoundFile(info.Name()) {
			parent := filepath.Dir(path)
			var cats []string
			for filepath.Base(parent) != filepath.Base(top) {
				c := filepath.Base(parent)
				if _, ok := catMap[c]; !ok {
					categories = append(categories, c)
					catMap[c] = true
				}
				cats = append(cats, c)
				parent = filepath.Dir(parent)
			}
			files = append(files, File{
				ID:         strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
				FilePath:   path,
				Categories: cats,
			})
		}
		return nil
	})
	return
}

func isAllLowercase(basename string) bool {
	for _, r := range basename {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isSoundFile(basename string) bool {
	var ext = filepath.Ext(basename)
	if len(ext) > 0 {
		ext = ext[1:] // trim the dot
	}
	return isAllLowercase(basename) && (ext == "mp3" || ext == "ogg" || ext == "m4a" || ext == "m4r" || ext == "wav" || ext == "dca")
}

func needsConversionToDCA(path string) bool {
	var ext = filepath.Ext(filepath.Base(path))
	if len(ext) > 0 {
		ext = ext[1:] // trim the dot
	}
	return ext != "dca"
}
