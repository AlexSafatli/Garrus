package sound

import (
	"os"
	"path/filepath"
	"strings"
)

type Library struct {
	RootPath string
	SoundMap map[string]File
}

func (l *Library) doConversions() []error {
	var errs []error
	for _, v := range l.SoundMap {
		if needsConversionToDCA(v.FilePath) {
			var newPath = strings.TrimSuffix(v.FilePath, filepath.Ext(v.FilePath)) + ".dca"
			err := saveSoundFileToDCA(v.FilePath, newPath)
			if err != nil {
				errs = append(errs, err)
			}
			v.FilePath = newPath
		}
	}
	return errs
}

func GetSounds(rootPath string) *Library {
	l := Library{RootPath: rootPath, SoundMap: make(map[string]File)}
	files, err := walkRootDirectory(rootPath, 0)
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
	return &l
}

func walkRootDirectory(root string, level uint) (files []File, err error) {
	var cleanRoot = filepath.Clean(root)
	err = filepath.Walk(cleanRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil || path == cleanRoot || strings.HasPrefix(info.Name(), ".") {
			return nil // ignore hidden files, etc.
		}
		if !info.IsDir() && isSoundFile(info.Name()) {
			var cats []string
			if level > 0 {
				cats = []string{filepath.Base(cleanRoot)}
			}
			files = append(files, File{
				ID:         strings.TrimSuffix(info.Name(), filepath.Ext(info.Name())),
				FilePath:   path,
				Categories: cats,
			})
		} else if info.IsDir() && cleanRoot != filepath.Dir(path) {
			var subfiles, err = walkRootDirectory(path, level+1)
			if err != nil {
				return err
			}
			files = append(files, subfiles...)
		} else if cleanRoot == filepath.Dir(path) {
			return nil
		}
		return filepath.SkipDir
	})
	return
}

func isSoundFile(basename string) bool {
	var ext = filepath.Ext(basename)
	if len(ext) > 0 {
		ext = ext[1:] // trim the dot
	}
	return ext == "mp3" || ext == "ogg" || ext == "m4a" || ext == "m4r" || ext == "wav" || ext == "dca"
}

func needsConversionToDCA(path string) bool {
	var ext = filepath.Ext(filepath.Base(path))
	if len(ext) > 0 {
		ext = ext[1:] // trim the dot
	}
	return ext != "dca"
}
