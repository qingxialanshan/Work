package global

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var Global_OS = runtime.GOOS

var gloabal_path string
var Cmd_win = os.Getenv("ComSpec")

func CopyFile(src, dst string) (int64, error) {
	sf, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	sfi, _ := sf.Stat()
	mode := sfi.Mode()
	if mode.IsDir() {
		//	fmt.Println("Directory")
		e := os.Rename(src, dst)
		return 2, e
	}
	defer sf.Close()
	df, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer df.Close()
	return io.Copy(df, sf)
}

func Exist(fname string) bool {
	_, e := os.Stat(fname)
	if e != nil {
		return false
	} else {
		return true
	}
}

func find(dir, pattern string) string {
	//fmt.Println("0000", dir)

	if Exist(dir) == false {
		return ""
	}
	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//fmt.Println("Reading " + dir)

	for _, file := range files {
		new_file := filepath.Join(dir, file.Name())

		if strings.Contains(file.Name(), pattern) {
			//fmt.Println("^^^", file.Name(), dir, pattern)
			gloabal_path = gloabal_path + "," + new_file
		}
		if file.IsDir() {

			find(new_file, pattern)
		}
		//if strings.Contains(file.Name,"-debug.apk")
	}
	return gloabal_path
}

//detect the device is T124 or T210
func Detect_devices(installdir string) bool {
	//if device is arch64 return 1, else return 0
	//device_arch, _ := exec.Command("bash", "-c", `. ~/.bashrc && adb shell getprop ro.product.cpu.abi`).Output()
	adb_path := filepath.Join(installdir, "android-sdk-linux", "platform-tools", "adb")
	device_arch, _ := exec.Command("bash", "-c", adb_path+" shell getprop ro.product.cpu.abi").Output()

	is_64 := false
	if strings.Contains(string(device_arch), "arm64-v8a") {
		is_64 = true
	} else if strings.Contains(string(device_arch), "armeabi-v7a") {
		is_64 = false
	} else {
		fmt.Println("cannot detect the device")
		os.Exit(2)
	}
	return is_64
}

func Redirector(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// deploy and install the samples' apk to device
func Deploy(installdir string) {

	out := find(filepath.Join(installdir, "Samples"), "-debug.apk")
	apks := strings.Split(out, ",")
	if Global_OS == "darwin" {
		Global_OS = "macosx"
	}
	adb_path := filepath.Join(installdir, "android-sdk-"+Global_OS, "platform-tools", "adb")
	var deploy_command, console string
	if Global_OS == "windows" {
		console = filepath.Join(installdir, "console.exe")
		deploy_command = "cmd /c echo Deploy Samples;cmd /c echo Waiting for device...;"
	} else {
		if Global_OS == "macosx" {
			console = filepath.Join(installdir, "console.app", "Contents", "MacOS", "console")
		} else {
			console = filepath.Join(installdir, "console")
		}

		deploy_command = "echo Deploy Samples;echo Waiting for device...;"
	}

	for i := 0; i < len(apks); i++ {

		if apks[i] == "" {
			continue
		}
		deploy_command = deploy_command + adb_path + " install -r " + apks[i] + ";"
	}

	deploy_log := filepath.Join(installdir, "_installer/deploy.log")

	//fmt.Println(deploy_log)
	var cuda_samples string
	gloabal_path = ""

	if Global_OS == "linux-x64" {
		if Detect_devices(installdir) {
			cuda_samples = find(filepath.Join(installdir, "CUDA_Samples", "7.0"), "-debug.apk")

		} else {
			cuda_samples = find(filepath.Join(installdir, "CUDA_Samples", "6.5"), "-debug.apk")

		}

		cudaapks := strings.Split(cuda_samples, ",")

		for i := 0; i < len(cudaapks); i++ {
			//deploying cuda samples
			if cudaapks[i] == "" {
				continue
			}
			deploy_command = deploy_command + adb_path + " install -r " + cudaapks[i] + ";"
		}
	}

	e := exec.Command(console, "-t", "Deploy Samples", "-l", deploy_log, deploy_command).Run()
	if e != nil {
		fmt.Println("Deploy Samples Failed.")
		os.Exit(2)
	}

	exec.Command(adb_path, "kill-server")
	if Global_OS == "windows" {
		Redirector(exec.Command("taskkill", "/F", "/IM", "adb.exe"))
	} else {
		cmd := exec.Command("/bin/sh", "-c", `ps -e|grep "adb"|awk '{print $1}'`)
		pid, err := cmd.Output()
		if err != nil {
			fmt.Println("failed")
		}

		pids := strings.Fields(string(pid[:]))

		if len(pids) > 0 {
			for i := 0; i < len(pids); i++ {
				exec.Command("kill", pids[i]).Start()
				//fmt.Println("Success")
			}
		}
	}
}

// deploy and install the samples' apk to device
func Deploy_v3a(installdir string) {

	out := find(filepath.Join(installdir, "Samples"), "-debug.apk")
	apks := strings.Split(out, ",")
	adb_path := filepath.Join(installdir, "android-sdk-"+Global_OS, "platform-tools", "adb")
	var deploy_command, console string
	if Global_OS == "windows" {
		console = filepath.Join(installdir, "console.exe")
		deploy_command = "cmd /c echo Deploy Samples;cmd /c echo Waiting for device...;"
	} else {
		if Global_OS == "macosx" {
			console = filepath.Join(installdir, "console.app", "Contents", "MacOS", "console")
		} else {
			console = filepath.Join(installdir, "console")
		}

		deploy_command = "echo Deploy Samples;echo Waiting for device...;"
	}

	for i := 0; i < len(apks); i++ {

		if apks[i] == "" {
			continue
		}
		deploy_command = deploy_command + adb_path + " install -r " + apks[i] + ";"
	}

	deploy_log := filepath.Join(installdir, "_installer/deploy.log")

	var cuda_samples string
	gloabal_path = ""

	cuda_samples = find(filepath.Join(installdir, "CUDA_Samples", "7.0"), "-debug-32.apk")

	cudaapks := strings.Split(cuda_samples, ",")

	for i := 0; i < len(cudaapks); i++ {
		//deploying cuda samples
		if cudaapks[i] == "" {
			continue
		}
		deploy_command = deploy_command + adb_path + " install -r " + cudaapks[i] + ";"
	}

	//fmt.Println(deploy_command)
	e := exec.Command(console, "-t", "Deploy Samples", "-l", deploy_log, deploy_command).Run()
	if e != nil {
		fmt.Println("Deploy Samples Failed.")
		os.Exit(2)
	}

	exec.Command(adb_path, "kill-server")
	if Global_OS == "windows" {
		Redirector(exec.Command("taskkill", "/F", "/IM", "adb.exe"))
	} else {
		cmd := exec.Command("/bin/sh", "-c", `ps -e|grep "adb"|awk '{print $1}'`)
		pid, err := cmd.Output()
		if err != nil {
			fmt.Println("failed")
		}

		pids := strings.Fields(string(pid[:]))

		if len(pids) > 0 {
			for i := 0; i < len(pids); i++ {
				exec.Command("kill", pids[i]).Start()
				fmt.Println("Success")
			}
		}
	}
}
