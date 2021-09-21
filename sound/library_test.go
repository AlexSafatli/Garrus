package sound

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const tmpTreeRoot = "testing"

func prepareTestRoot() (string, error) {
	tmp, err := ioutil.TempDir("", tmpTreeRoot)
	if err != nil {
		return "", fmt.Errorf("Could not create temporary dir: %v\n", err)
	}
	return tmp, err
}

func prepareTestTree(tmp, tree string) (string, error) {
	var path = filepath.Join(tmp, tree)
	err := os.MkdirAll(path, 0755)
	if err != nil {
		_ = os.RemoveAll(path)
		return "", err
	}
	empty, err := os.Create(filepath.Join(path, "empty.mp3"))
	if err != nil {
		_ = os.RemoveAll(path)
		return "", err
	}
	_ = empty.Close()
	return path, err
}

func TestWalkRootDirectoryForSounds(t *testing.T) {
	tmp, err := prepareTestRoot()
	if err != nil {
		t.Error(err)
		return
	}

	tmpA, errA := prepareTestTree(tmp, "Games/Factorio")
	tmpB, errB := prepareTestTree(tmp, "")
	tmpC, errC := prepareTestTree(tmp, "Music")
	if errA != nil {
		t.Error(errA)
	}
	if errB != nil {
		t.Error(errB)
	}
	if errC != nil {
		t.Error(errC)
	}
	defer os.RemoveAll(tmpA)
	defer os.RemoveAll(tmpB)
	defer os.RemoveAll(tmpC)

	f, err := walkRootDirectoryForSounds(tmp, "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(f) != 3 {
		t.Errorf("Root: %s\nDid not find three paths, found %d instead\n%+v",
			tmp, len(f), f)
		return
	}
	if f[0].ID != "empty" || f[1].ID != "empty" || f[2].ID != "empty" {
		t.Errorf("%+v contained incorrect ID data", f)
	}
	if len(f[0].Categories) != 2 || f[0].Categories[0] != "Factorio" || f[0].Categories[1] != "Games" {
		t.Errorf("%+v contained incorrect category data", f[0])
	}
	if len(f[1].Categories) == 0 || f[1].Categories[0] != "Music" {
		t.Errorf("%+v contained incorrect category data", f[1])
	}
	if len(f[2].Categories) > 0 {
		t.Errorf("%+v contained incorrect category data", f[2])
	}
}

func TestWalkRootDirectoryForCategories(t *testing.T) {
	tmp, err := prepareTestRoot()
	if err != nil {
		t.Error(err)
		return
	}

	tmpA, errA := prepareTestTree(tmp, "Games/Factorio")
	tmpB, errB := prepareTestTree(tmp, "")
	tmpC, errC := prepareTestTree(tmp, "Music")
	tmpD, errD := prepareTestTree(tmp, "TV Series/Star Trek")
	if errA != nil {
		t.Error(errA)
	}
	if errB != nil {
		t.Error(errB)
	}
	if errC != nil {
		t.Error(errC)
	}
	if errD != nil {
		t.Error(errC)
	}
	defer os.RemoveAll(tmpA)
	defer os.RemoveAll(tmpB)
	defer os.RemoveAll(tmpC)
	defer os.RemoveAll(tmpD)

	f, err := walkRootDirectoryForCategories(tmp, "")
	if err != nil {
		t.Error(err)
		return
	}
	if len(f) != 3 {
		t.Errorf("Root: %s\nDid not find three categories, found %d instead\n%+v",
			tmp, len(f), f)
		return
	}
	if f[0].Name != "Games" && len(f[0].Children) != 1 {
		t.Errorf("%+v contained incorrect ID data", f[0])
	}
	if f[1].Name != "Music" && len(f[0].Children) != 0 {
		t.Errorf("%+v contained incorrect ID data", f[1])
	}
	if f[2].Name != "TV Series" && len(f[0].Children) != 1 {
		t.Errorf("%+v contained incorrect ID data", f[2])
	}
}
