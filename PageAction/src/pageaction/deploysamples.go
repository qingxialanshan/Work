package pageaction

import (
	"fmt"
	"global"
	"os"
)

type DeploySamples struct {
	PageAction
}

func (deploy *DeploySamples) PostAction(args ...string) {
	//PageAction deploy_samples next_page
	nextpage := args[0]

	if nextpage == "nextpage" {
		deploy.NextPage = "finished"
		fmt.Println(deploy.NextPage)

	} else if nextpage == "install" {
		if len(args) < 3 {
			fmt.Println(`Wrong args input. Please input as followings
			PageAction deploy_samples install [installdir] [version]`)
			os.Exit(-1)
		}
		installdir := args[1]
		version := args[2]
		//fmt.Println(installdir, version)
		if version == "1.0" {
			global.Deploy_v3a(installdir)
		} else {
			//fmt.Println("deploy ...")
			global.Deploy(installdir)
		}
	} else {
		fmt.Println("Wrong args input.")
		os.Exit(-1)
	}

}
