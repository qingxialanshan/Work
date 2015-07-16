package global

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type args struct {
	cmd  string
	args string
}

var (
	osCommand = map[string]*args{
		"darwin":  &args{"open", ""},
		"freebsd": &args{"xdg-open", ""},
		"linux":   &args{"xdg-open", ""},
		"netbsd":  &args{"xdg-open", ""},
		"openbsd": &args{"xdg-open", ""}, // It may be open instead
		"windows": &args{os.Getenv("ComSpec"), "/c start "},
	}
	ErrCantOpen     = errors.New("webbrowser.Open: can't open webpage")
	ErrNoCandidates = errors.New("webbrowser.Open: no browser candidate found for your OS.")
)

func OpenDocs(s string) (e error) {
	if os, ok := osCommand[runtime.GOOS]; ok {
		//fmt.Println(os.cmd, os.args, s)

		_, e = exec.Command(os.cmd, os.args+s).Output()
		if e != nil {
			fmt.Println(e)
		}
	}
	return
}
