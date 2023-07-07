package server

import (
	"archive/zip"
	"fmt"
	"io"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/ghostx31/nativefier-downloader/internal/structs"
)

// GetUrlFromUser function is responsible for collecting URL, OS and other data from the user.
func GetUrlFromUser(urlparams structs.Urlparams) string {
	fileName := BuildWebApp(urlparams)
	return fileName
}

// GetFilename is responsible for getting the name of the zip file and the directory created.
func GetFilename(urlparams structs.Urlparams) (zipFileName string, folderName string, directoryName string) {
	name, err := url.Parse(urlparams.Url)
	if err != nil {
		panic(err)
	}

	if urlparams.Os == "windows" {
		urlparams.Os = "win32"
	}
	if urlparams.Os == "mac" {
		urlparams.Os = "darwin"
	}

	folderName = name.Hostname()
	runtimeOs := strings.ReplaceAll(runtime.GOARCH, "amd", "x")
	directoryName = folderName + "-" + urlparams.Os + "-" + runtimeOs
	zipFileName = directoryName + ".zip"

	fmt.Printf("\nDirectory name is: %v", directoryName)
	fmt.Printf("\nZip file name is: %v", zipFileName)
	file, err := os.OpenFile("save.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	_, err = file.Write([]byte(zipFileName))
	if err != nil {
		panic(err)
	}
	return zipFileName, folderName, directoryName
}

// BuildWebApp is responsible for building the electron app.
func BuildWebApp(urlparams structs.Urlparams) string {
	zipFileName, folderName, directoryName := GetFilename(urlparams)
	executeCommand := exec.Command("./node_modules/.bin/nativefier", urlparams.Url, "--name", folderName, "-p", urlparams.Os, "--tray", urlparams.Tray, "--widevine", urlparams.Widevine)
	return execCommandChore(executeCommand, zipFileName, directoryName)
}

// zipDirectory creates a zip of the electron app directory built by BuildWebApp
func zipDirectory(source, target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	writer := zip.NewWriter(f)
	defer func(writer *zip.Writer) {
		err := writer.Close()
		if err != nil {
			panic(err)
		}
	}(writer)

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate

		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		}
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}(f)

		_, err = io.Copy(headerWriter, f)
		return err
	})
}

func isRunningInDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}
	return false
}

func execCommandChore(executeCommand *exec.Cmd, zipFileName string, directoryName string) string {
	stdout, err := executeCommand.Output()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf(string(stdout))
	fmt.Printf("Zipping: %v\n", zipFileName)
	err = zipDirectory(directoryName, zipFileName)
	fmt.Printf("Zip complete! %v\n", zipFileName)
	return zipFileName
}
