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

func TestWalkRootDirectory(t *testing.T) {
	tmp, err := prepareTestRoot()
	if err != nil {
		t.Error(err)
		return
	}

	tmpA, errA := prepareTestTree(tmp, "Games/Factorio")
	tmpB, errB := prepareTestTree(tmp, "Games/Starcraft")
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

	f, err := walkRootDirectory(tmp, 0)
	if err != nil {
		t.Error(err)
		return
	}
	if len(f) != 3 {
		t.Errorf("Root: %s\nDid not find three paths, found %d instead\n%+v",
			os.TempDir(), len(f), f)
		return
	}
	if f[0].ID != "empty" || f[1].ID != "empty" || f[2].ID != "empty" {
		t.Errorf("%+v contained incorrect data", f)
	}
}
