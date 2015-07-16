//args[0]: PageAction
//args[1]: finished
//args[2]: create_workspace|create_shortcut|browse_docs|restart|remove_logs|browse_ibdocs

package pageaction

import (
	"global"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
    "fmt"    
)

type Finished struct {
	PageAction
}

func (f *Finished) PostAction(args ...string) {
	//finished page actions
	//option
	option := args[0]

	installdir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	if option == "create_shortcut" {
		if global.Global_OS == "windows" {
			win_sc_folder := filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Windows", "Start Menu", "Programs", "NIDIA Corporation")
			cmd := os.Getenv("ComSpec")
			if global.ComponentSelect("docs") {
				exec.Command(cmd, "/c", `mklink "AndroidWorks Docs" `+filepath.Join(installdir, "docs", "index.html")).Output()
				os.Rename("AndroidWorks Docs", win_sc_folder)
			}

		}
		return
	}

	if option == "remove_logs" {
		os.RemoveAll(filepath.Join(installdir, "_installer"))
	}

	if option == "browse_docs" {
		var url string
		if global.ComponentSelect("docs") {
			installdir = strings.Replace(installdir, "\\", "/", -1)
			//url = "file://" + installdir + "/docs/index.html"
			url = filepath.Join(installdir, "docs", "index.html")

		} else {
			url = "http://docs.nvidia.com/gameworks/index.html#developertools/mobile/androidworks/androidworks.htm"
		}
		global.OpenDocs(url)
		return
	}
	if option == "restart" {
		exec.Command(os.Getenv("ComSpec"), "/c", "shutdown /r /f").Output()
		return
	}

	if option == "browse_ibdocs" {
		//browse the incredibuild docs (incredibuild.chm)
		e1 := exec.Command("cmd", "/c", filepath.Join(os.Getenv("ProgramFiles(x86)"), "Xoreax", "IncrediBuild", "IncrediBuild.chm")).Start()
        if e1!=nil{
            fmt.Println(e1)
            os.Exit(2)
        }
		
	}

	if option == "create_workspace" {
		var create_workspace = 0
		//create the nvsample workspace for eclipse
		if global.ComponentSelect("eclipse") {
			if global.ComponentSelect("tdksample") || global.ComponentSelect("opencv") ||
				global.ComponentSelect("cudasamples65") || global.ComponentSelect("cudasamples70") {
				create_workspace = 1
			}
		}
		//fmt.Println(create_workspace)
		if create_workspace == 1 {
			var eclipse string
			if global.Global_OS == "windows" {
				eclipse = filepath.Join(installdir, "eclipse", "eclipse.exe")
			} else {
				eclipse = filepath.Join(installdir, "eclipse", "eclipse")
			}
			if global.ComponentSelect("tdksample") {
				exec.Command(eclipse, "-nosplash", "-data", filepath.Join(installdir, "nvsample_workspace"),
					"-application", "org.eclipse.cdt.managedbuilder.core.headlessbuild", "-importAll",
					filepath.Join(installdir, "Samples", "TDK_Samples")).Output()
				//fmt.Println(e)
			}
			if global.ComponentSelect("opencv") {
				exec.Command(eclipse, "-nosplash", "-data", filepath.Join(installdir, "nvsample_workspace"),
					"-application", "org.eclipse.cdt.managedbuilder.core.headlessbuild", "-importAll",
					filepath.Join(installdir, "OpenCV-2.4.8.2-Tegra-sdk")).Output()
			}
			if global.ComponentSelect("cudasamples65") || global.ComponentSelect("cudasamples70") {
				var cudasample_path = filepath.Join(installdir, "CUDA_Samples")
				if global.Exist(filepath.Join(installdir, "CUDA_Samples", "7.0")) {
					cudasample_path = filepath.Join(cudasample_path, "7.0")
				} else {
					cudasample_path = cudasample_path
				}

				exec.Command(eclipse, "-nosplash", "-data", filepath.Join(installdir, "nvsample_workspace"),
					"-application", "org.eclipse.cdt.managedbuilder.core.headlessbuild", "-importAll",
					filepath.Join(installdir, "Samples", cudasample_path)).Output()
			}
		}
	}
}
