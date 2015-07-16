package pageaction

import (
	"fmt"
)

type InstallConfig struct {
	PageAction
}

func (ic *InstallConfig) PostAction(args ...string) {
	//	if len(args) < 1 {
	//		fmt.Println("Please input the installdir")
	//		os.Exit(-1)
	//	}
	//	installdir := args[0]
	//	os.Mkdir(installdir, 0777)
	//copy the files to installdir on windows
	//	if global.Global_OS == "windows" {
	//		global.CopyFile(args[0], args[1])
	//	}
	ic.NextPage = "installation"
	fmt.Println(ic.NextPage)
}
