package sound

import (
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Library struct {
	RootPath   string
	SoundMap   map[string]File
	Categories []Category
}

func (l *Library) doConversions() []error {
	var errs []error
	for _, v := range l.SoundMap {
		if needsConversionToDCA(v.FilePath) {
			var newPath = strings.TrimSuffix(v.FilePath, filepath.Ext(v.FilePath)) + ".dca"
			err := saveSoundFileToDCA(v.FilePath, newPath)
			if err != nil {
				errs = append(errs, err)
			} else {
				err = os.Remove(v.FilePath)
				if err != nil {
					errs = append(errs, err)
				}
			}
			v.FilePath = newPath
		}
	}
	return errs
}

func GetSounds(rootPath string) *Library {
	l := Library{RootPath: rootPath, SoundMap: make(map[string]File)}
	files, err := walkRootDirectoryForSounds(rootPath, "")
	if err != nil {
		return nil
	}
	for i := range files {
		l.SoundMap[files[i].ID] = files[i]
	}
	errs := l.doConversions()
	if len(errs) > 0 {
		return nil
	}
	cats, err := walkRootDirectoryForCategories(rootPath, "")
	if err != nil {
		return nil
	}
	l.Categories = cats
	return &l
}

func walkRootDirectoryForSounds(start, root string) (files []File, err error) {
	if len(root) == 0 {
		root = start
	}
	var cleanRoot = filepath.Clean(start)
	err = filepath.Walk(cleanRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil || path == cleanRoot || strings.HasPrefix(info.Name(), ".") {
			return nil // ignore hidden files, etc.
		}
		if !info.IsDir() && isSoundFile(info.Name()) {
			var cats []string
			if filepath.Dir(path) != root {
				cats = []string{filepath.Base(filepath.Dir(path))}
			}
			files = append(files, File{
				ID:         strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
				FilePath:   path,
				Categories: cats,
			})
		} else if info.IsDir() && cleanRoot != filepath.Dir(path) {
			var subfiles, err = walkRootDirectoryForSounds(path, root)
			if err != nil {
				return err
			}
			for i := range subfiles {
				subfiles[i].Categories = append(subfiles[i].Categories, filepath.Base(filepath.Dir(path)))
			}
			files = append(files, subfiles...)
		} else if cleanRoot == filepath.Dir(path) {
			return nil
		}
		return filepath.SkipDir
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
