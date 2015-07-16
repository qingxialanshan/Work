package pageaction

import (
	"fmt"
	"global"
	"os"
	"os/exec"
	"path/filepath"
)

type Installation struct {
	PageAction
}

func (i *Installation) PostAction(args ...string) {
	if len(args) < 1 {
		fmt.Println("Please input the option nextpage|install to continue")
		os.Exit(-1)
	}
	option := args[0]

	if option == "nextpage" {
		i.NextPage = "finished"

		if global.ComponentSelect("drivertools") {
			i.NextPage = "device_setup"
		} else if global.ComponentSelect("allowcompile") {
			i.NextPage = "deploy_samples"
		} else {
			i.NextPage = "finished"
		}
		fmt.Println(i.NextPage)
	} else if option == "install" {
		//copy the files to the installdir
		if len(args) < 4 {
			fmt.Println("Please input the args as followings ")
			fmt.Println("PageAction installation install [pksrc] [installdir] [exefile]")
			os.Exit(2)
		}
		pksrc := args[1]
		installdir := args[2]
		exefile := args[3]

		if _, e1 := os.Stat(installdir); e1 != nil {
			os.Mkdir(installdir, 0777)
		}

		if _, se := os.Stat(filepath.Join(installdir, "_installer")); se != nil {
			os.Mkdir(filepath.Join(installdir, "_installer"), 0777)
		}
		e := global.CopyFileExp(pksrc, installdir, exefile)
		if e != nil {
			fmt.Println(e)
		}
		//fmt.Println(os.Args[0])

		if global.Global_OS == "darwin" {
			global.CopyFile(os.Args[0], filepath.Join(installdir, "PageAction"))
			os.Chmod(filepath.Join(installdir, "PageAction"), 0777)
		} else if global.Global_OS == "linux" {
			//add the Chooser to desktop entry
			chooser_entry := `[Desktop Entry]
Type=Application
Name=AndroidWorks Component Manager
GenericName=AndroidWorks Component Manager
Icon=` + filepath.Join(installdir, "TegraDeveloperKit.ico") + `
Exec=` + filepath.Join(installdir, "Chooser") + `
TryExec=` + filepath.Join(installdir, "Chooser") + `
Keywords=androidworks;nvidia;
Terminal=No
Categories=Development;`
			entry_folder := os.Getenv("HOME") + "/.local/share/applications"
			if _, se := os.Stat(entry_folder); se != nil {
				//fmt.Println("no folder exist")
				_, ce := exec.Command("mkdir", "-p", entry_folder).Output()
				if ce != nil {
					fmt.Println(ce)
				}
			}
			fp, fe := os.OpenFile(entry_folder+"/AndroidWorks.desktop", os.O_CREATE|os.O_RDWR, 0600)
			if fe != nil {
				fmt.Println(fe)
			}
			fp.WriteString(chooser_entry)
		}
		//set the uninstaller to the registery
		if global.Global_OS == "windows" {
			_, e0 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/v", "DisplayVersion", "/d", `1R1`, "/f").Output()
			_, e1 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/v", "NoModify", "/d", "0", "/f").Output()
			_, e2 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/v", "Publisher", "/d", `NVIDIA Corpration`, "/f").Output()
			_, e3 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/v", "DisplayIcon", "/d", filepath.Join(installdir, "Chooser.exe"), "/f").Output()
			_, e4 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/v", "ModifyPath", "/d", filepath.Join(installdir, "Chooser.exe"), "/f").Output()
			_, e5 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/v", "UninstallString", "/d", filepath.Join(installdir, "Uninstaller.bat"), "/f").Output()
			_, e6 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\microsoft\windows\CurrentVersion\Uninstall\NVIDIA AndroidWorks`, "/v", "DisplayName", "/d", `NVIDIA AndroidWorks`, "/f").Output()

			if e0 != nil || e1 != nil || e2 != nil || e3 != nil || e4 != nil || e5 != nil || e6 != nil {
				fmt.Println(e1, e2, e3, e4, e5, e6)
			}

			df, _ := os.OpenFile(filepath.Join(installdir, "uninstaller.bat"), os.O_CREATE|os.O_RDWR, 0777)
			_, e7 := df.WriteString("@cd %~dp0\r\n@set /a count=0\r\n@start /wait Chooser.exe -u -v\r\n" +
				"@ping 127.0.0.1 -n 10 -w 1000>null\r\n" + ":Repeat\r\n" + "@del Chooser.exe\r\n" + "@set /a count=%count%+1\r\n" +
				"@ping 127.0.0.1 -n 3 -w 1000>null\r\n" + "@if exist Chooser.exe (\r\n" + "@if %count NEQ 60 goto Repeat\r\n" + ")\r\n" +
				"@reg delete \"HKLM\\Software\\Wow6432Node\\microsoft\\windows\\CurrentVersion\\Uninstall\\NVIDIA AndroidWorks\" /f\r\n" +
				"@rmdir /S /Q " + installdir + "\r\n@del null\r\n" + "@del Uninstaller.bat\r\n")
			if e7 != nil {
				fmt.Println(e7)
			}
			//add the AndroidWorks to HKLM\Software\Wow6432Node\NVIDIA Corporation\AndroidWorks
			_, e8 := exec.Command("reg", "add", `HKLM\Software\Wow6432Node\NVIDIA Corporation\AndroidWorks`, "/v", "Location", "/d", installdir, "/f").Output()
			if e8 != nil {
				fmt.Println(e8)
			}
		}
	}

}
