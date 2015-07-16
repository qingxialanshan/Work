package pageaction

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

import (
	"global"
)

type InstallDependency struct {
	PageAction
}

func (id *InstallDependency) PostAction(args ...string) {
	if global.Global_OS == "linux" {
		linux_version, _ := exec.Command("cat", "/etc/issue").Output()
		linux_ver := string(linux_version[:])
		if strings.Contains(linux_ver, "12.04") {
			linux_ver = "12.04"
		} else if strings.Contains(linux_ver, "14.04") {
			linux_ver = "14.04"
		} else {
			linux_ver = ""
		}
		//fmt.Println(linux_ver)
		//		cmd := ""

		_, dependpkg := CheckDependency()
		//		if strings.Contains(dependpkg, "ia32-libs") {

		//			//			_, e := exec.Command("bash", "-c", "xterm -T 'Installing OS Dependency on 12.04' -e sudo apt-get install ia32-libs").Output()

		//			//			if e != nil {
		//			//				fmt.Print(e)
		//			//				os.Exit(2)
		//			//			}
		//			cmd = "-T 'Installing OS Dependency on 12.04' -e sudo apt-get install -y ia32-libs"
		//		} else if strings.Contains(dependpkg, "libc6:i386") {
		//			//			_, e := exec.Command("bash", "-c", "xterm -T 'Installing OS Dependency on 14.04' -e sudo dpkg --add-architecture i386 && "+
		//			//				"sudo apt-get update && sudo -s apt-get install libc6:i386 libncurses5:i386 libstdc++6:i386 lib32z1 lib32ncurses5 "+
		//			//				"lib32bz2-1.0bs").Output()
		//			//			if e != nil {
		//			//				fmt.Print(e)
		//			//				os.Exit(2)
		//			//			}
		//			cmd = "-T 'Installing OS Dependency on 14.04' -e sudo dpkg --add-architecture i386 && " +
		//				"sudo apt-get update && sudo -s apt-get install -y libc6:i386 libncurses5:i386 libstdc++6:i386 lib32z1 lib32ncurses5 " +
		//				"lib32bz2-1.0bs"
		//		}
		//		if strings.Contains(dependpkg, "dos2unix") {
		//			if cmd != "" {
		//				cmd = cmd + " dos2unix"
		//			} else {
		//				cmd = cmd + "-e sudo -s apt-get install dos2unix"
		//			}
		//		}
		//		if strings.Contains(dependpkg, "openjdk-7-jdk") {
		//			if cmd != "" {
		//				cmd = cmd + " openjdk-7-jdk"
		//			} else {
		//				cmd = cmd + "-e sudo -s apt-get install -y openjdk-7-jdk"
		//			}
		//		}
		if dependpkg != "" {
			dependpkg = strings.Replace(dependpkg, ",", " ", -1)
			_, e := exec.Command("bash", "-c", "xterm -e sudo -s apt-get install -y "+dependpkg).Output()
			if e != nil {
				fmt.Println(e)
				os.Exit(2)
			}
		}
	}
	id.NextPage = "install_config"
	fmt.Println(id.NextPage)
}
