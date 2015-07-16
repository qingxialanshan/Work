package global

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFileExp(t *testing.T) {
	src := "/home/amyl/Downloads/test"
	des := "/home/amyl/Downloads/test1"
	exfile := "dmg2img.exe"
	err := CopyFileExp(src, des, exfile)
	if err != nil {
		t.Fatalf("CopyFileExp %s %s %s:%v", src, des, exfile, err)
	}

	foundFile := true
	foundExp := false

	list, err := ioutil.ReadDir(src)
	for _, dir := range list {
		if dir.Name() == exfile {
			if _, e := os.Stat(filepath.Join(des, exfile)); e == nil {
				foundExp = true
				break
			}

		}

	}
	for _, dir := range list {

		if _, e := os.Stat(filepath.Join(des, dir.Name())); e != nil && dir.Name() != exfile {
			foundFile = false
			break
		}
	}
	if !foundFile {
		t.Fatalf("CopyFileExp %s %s %s: not found", src, des, exfile)
	}
	if foundExp {
		t.Fatalf("CopyFileExp %s found", exfile)
	}
}
