//uninstaller for tadp
//1. call Chooser -u -i to uninstall all the installed components
//2. remove files according to .dat
//3. delete shortcut
//4. remove registry

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	//step 1: Chooser -u -i
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	var chooser string
	fmt.Println(runtime.GOOS)
	os.Chdir(dir)
	if runtime.GOOS == "windows" {
		chooser = filepath.Join(dir, "Chooser.exe")
	} else {
		chooser = filepath.Join(dir, "Chooser")
	}

	if runtime.GOOS == "windows" {
		exec.Command("sudo", chooser, "-u", "-v").Run()

	} else {
		fmt.Println("entering...")
		e := exec.Command(chooser, "-u", "-v").Run()
		if e != nil {
			os.Exit(-1)
		}

		fmt.Println(os.Args[0])
		e = os.Remove(os.Args[0])
		if e != nil {
			fmt.Println("the error is:", e)
			os.Exit(2)
		}

	}

	//step 2: remove files from .dat
	dat_file := filepath.Join(os.Getenv("HOME"), ".res.dat")
	if runtime.GOOS == "windows" {
		dat_file = filepath.Join(os.Getenv("USERPROFILE"), ".res.dat")
	}

	inf, _ := os.Open(dat_file)
	fd, _ := ioutil.ReadAll(inf)

	files := strings.Split(string(fd), ",")
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
		if e := os.Remove(files[i]); e != nil {
			fmt.Println(e)
		}
	}

	//3. delete shortcut

	if runtime.GOOS == "windows" {
		fmt.Println("remove registry")
		//step 3: delete shortcut
		win_sc_folder := filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "NIDIA Corporation")
		os.Remove(filepath.Join(win_sc_folder, "AndroidWorks docs"))

		//step 4: remove registry
		reg := filepath.Join(os.Getenv("SystemRoot"), "system32", "reg.exe")
		out, e := exec.Command(reg, "delete", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/f").Output()
		if e != nil {
			fmt.Println(e, out)
		}

	}

	//4. delete itself
	if runtime.GOOS != "windows" {
		exec.Command("rm", "-rf", dir).Run()
	}
}
