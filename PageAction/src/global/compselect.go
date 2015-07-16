package global

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var compFile string

func ComponentSelect(compName string) bool {
	if Global_OS == "windows" {
		compFile = filepath.Join(os.Getenv("USERPROFILE"), "selectcomp.txt")
	} else {
		compFile = filepath.Join(os.Getenv("HOME"), "selectcomp.txt")
	}
	inf, e := os.Open(compFile)

	fd, err := ioutil.ReadAll(inf)
	if err != nil || e != nil {
		//fmt.Println("read file failed")
	}
	complist := string(fd)
	//fmt.Println(complist)
	if strings.Contains(complist, compName) {
		return true
	}
	return false
}
