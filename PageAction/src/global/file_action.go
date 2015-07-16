package global

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFileExp(src, dest, efile string) error {
	//copy files from the src to dest except the efile
	if _, e := os.Stat(src); e != nil {
		return e
	}
	if df, e := os.Stat(dest); e != nil || df.IsDir() == false {

		err := os.Mkdir(dest, 0777)
		if err != nil {
			fmt.Println("creat error")
		}
	}

	s, oe := os.Open(src)
	if oe != nil {
		return oe
	}
	defer s.Close()

	files, re := s.Readdir(-1)
	if re != nil {
		return re
	}

	var write_to_dat string
	done := 0
	for _, file := range files {
		if file.Name() == efile {
			continue
		}
		//fmt.Println(file.Name())
		sourcefilepointer := filepath.Join(src, file.Name())

		destinationfilepointer := filepath.Join(dest, file.Name())
		var e error
		if file.IsDir() {
			e = copyDir(sourcefilepointer, destinationfilepointer)
		} else {
			e = copyFile(sourcefilepointer, destinationfilepointer)
		}
		if e != nil {
			return e
		}

	}
	if done == 1 {
		return errors.New("Copy Failed.")
	}
	//fmt.Println(write_to_dat)
	//put the write_to_dat to .res.dat
	if write_to_dat != "" {
		dat, _ := os.OpenFile(".res.dat", os.O_CREATE|os.O_RDWR, 0600)
		defer dat.Close()
		dat.WriteString(write_to_dat)

	}

	return nil
}

func copyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	if _, e := os.Stat(dest); e == nil {
		//fmt.Println("the file is not overide able", filepath.Join(dest, file.Name()))
		os.Chmod(dest, 0777)
	}

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		fmt.Println(sourceinfo.Mode())
		if err == nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func copyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = copyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = copyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}
