package sound

import (
	"github.com/AlexSafatli/Garrus/structs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Library struct {
	RootPath   string
	SoundMap   map[string]*File
	Categories []Category
	Trie       *structs.LowercaseTrie
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

func (l *Library) GetSoundNames() []string {
	keys := make([]string, 0, len(l.SoundMap))
	for k := range l.SoundMap {
		keys = append(keys, k)
	}
	return keys
}

func (l *Library) Contains(s string) bool {
	return l.Trie.Contains(s)
}

func (l *Library) GetClosestMatchingSoundID(s string) string {
	return l.Trie.GetWordWithPrefix(s)
}

func LoadSounds(rootPath string) error {
	l := Library{RootPath: rootPath, SoundMap: make(map[string]*File)}
	files, err := walkRootDirectoryForSounds(rootPath, "")
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
	cats, err := walkRootDirectoryForCategories(rootPath, "")
	if err != nil {
		return err
	}
	l.Categories = cats
	l.Trie = structs.NewLowercaseTrie(l.GetSoundNames())
	library = l
	return nil
}

func GetLibrary() *Library {
	return &library
}

func walkRootDirectoryForSounds(start, top string) (files []File, err error) {
	var cleanStart = filepath.Clean(start)
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
				cats = append(cats, filepath.Base(parent))
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

func walkRootDirectoryForCategories(start, root string) (cats []Category, err error) {
	if len(root) == 0 {
		root = start
	}
	var cleanRoot = filepath.Clean(start)
	err = filepath.Walk(cleanRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil || path == cleanRoot || strings.HasPrefix(info.Name(), ".") {
			return nil // ignore hidden files, etc.
		}
		if info.IsDir() {
			var cat = Category{Name: info.Name()}
			children, err := walkRootDirectoryForCategories(path, root)
			if err != nil {
				return err
			}
			cat.Children = append(cat.Children, children...)
			cats = append(cats, cat)
		}
		return filepath.SkipDir
	})
	return
}
