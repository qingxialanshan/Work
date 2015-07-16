package pageaction

import (
	"fmt"
	"os/exec"
	"strings"
)

import (
	"global"
)

type Suggestion struct {
	PageAction
}

func (sug *Suggestion) PostAction(args ...string) {
	//fmt.Println(CheckDependency())
	need_depend, depend_pkg := CheckDependency()

	if (global.Global_OS == "linux") && (need_depend) {
		sug.NextPage = "install_dependency"
		fmt.Println(sug.NextPage + ":" + depend_pkg)

	} else {
		sug.NextPage = "install_config"
		fmt.Println(sug.NextPage)
	}

}

func CheckDependency() (need_dependency bool, depend_pkg string) {
	//check the dependency for linux-x64
	var lib_32 string
	need_dependency = false
	linux_version, _ := exec.Command("cat", "/etc/issue").Output()
	linux_ver := string(linux_version[:])

	if strings.Contains(linux_ver, "12.04") {
		//if the linux is Ubuntu 12.04 then check the ia32-libs
		if _, err := exec.Command("dpkg", "-s", "ia32-libs").Output(); err != nil {

			lib_32 = "ia32-libs"
			need_dependency = true
		}

	} else if strings.Contains(linux_ver, "14.04") {
		lib_32s := "lib32z1,lib32ncurses5,lib32bz2-1.0,libc6:i386,libncurses5:i386,libstdc++6:i386"

		for i := 0; i < len(strings.Split(lib_32s, ",")); i++ {
			if _, err := exec.Command("dpkg", "-s", strings.Split(lib_32s, ",")[i]).Output(); err != nil {
				need_dependency = true
				if lib_32 == "" {
					lib_32 = strings.Split(lib_32s, ",")[i]
				} else {
					lib_32 = lib_32 + "," + strings.Split(lib_32s, ",")[i]
				}
			}
		}

	}

	if _, err1 := exec.Command("dpkg", "-s", "dos2unix").Output(); err1 != nil {
		need_dependency = true
		if lib_32 != "" {
			lib_32 = lib_32 + ",dos2unix"
		} else {
			lib_32 = "dos2unix"
		}
	}
	if _, err2 := exec.Command("dpkg", "-s", "openjdk-7-jdk").Output(); err2 != nil {
		need_dependency = true
		if lib_32 != "" {
			lib_32 = lib_32 + ",openjdk-7-jdk"
		} else {
			lib_32 = "openjdk-7-jdk"
		}
	}
	depend_pkg = lib_32

	return
}
