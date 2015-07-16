package pageaction

import (
	"fmt"
	"global"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type DeviceSetup struct {
	PageAction
}

func (ds *DeviceSetup) PostAction(args ...string) {
	//device_setup nextpage|install
	action := args[0]

	if action == "nextpage" {
		//set the NextPage
		if global.ComponentSelect("allowcompile") {
			ds.NextPage = "deploy_samples"
		} else {
			ds.NextPage = "finished"
		}
		fmt.Println(ds.NextPage)
	} else if action == "install" {
		if len(args) < 2 {
			fmt.Println(`Wrong args input. Please input as followings
			PageAction device_setup install [installdir]`)
			os.Exit(-1)
		}
		installdir := args[1]
		adb_path := filepath.Join(installdir, "android-sdk-linux", "platform-tools", "adb")
		console := filepath.Join(installdir, "console")

		//deploy the cuda runtime to devices
		//fmt.Println("the devices have been detected: ", global.Detect_devices(installdir))
		if global.Detect_devices(installdir) {
			//device is arm64-v8a
			exec.Command(adb_path, "wait-for-device").Output()
			oe, re := exec.Command(adb_path, "root").Output()
			if re != nil || strings.Contains(string(oe), "adbd cannot run as root in production builds") {
				//adb root error
				fmt.Println("Your device is not connected or not allowed adb root.")
				os.Exit(2)
			}
			cuda_runtime := filepath.Join(installdir, "cuda-android-7.0", "aarch64-linux-androideabi")
			cuda_device := "/data/cuda-toolkit-7.0"
			_, e := exec.Command(console, "-t", "Pushing cuda-toolkit-7.0", adb_path+" push "+cuda_runtime+" "+cuda_device).Output()
			if e != nil {
				fmt.Print(e)
				os.Exit(2)
			}
		} else {
			//device is armeabi-v7a
			exec.Command(adb_path, "wait-for-device").Output()
			oe, re := exec.Command(adb_path, "root").Output()
			if re != nil || strings.Contains(string(oe), "adbd cannot run as root in production builds") {
				//adb root error
				fmt.Println("Your device is not connected or not allowed adb root.")
				os.Exit(2)
			}
			cuda_runtime := filepath.Join(installdir, "cuda-android-6.5")
			cuda_device := "/data/cuda-toolkit-6.0"
			_, e := exec.Command(console, "-t", "Pushing cuda-toolkit-6.5", adb_path+" push "+cuda_runtime+" "+cuda_device).Output()
			if e != nil {
				fmt.Print(e)
				os.Exit(2)
			}
		}
	} else {
		fmt.Println("Wrong args input.")
		os.Exit(-1)
	}

}
