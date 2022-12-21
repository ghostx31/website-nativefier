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
)

func GetUrlFromUser(Url string, Os string) string {
	fileName := BuildWebApp(Url, Os)

	return fileName
}

func BuildWebApp(Url string, Os string) string {
	name, err := url.Parse(Url)
	if err != nil {
		panic(err)
	}

	if Os == "windows" { // If OS is Windows
		Os = "win32"
	}
	if Os == "mac" { // If OS is Mac
		Os = "darwin"
	}
	// TODO: Add support for more mac related stuff. nativefier --help for more info.
	folderName := name.Hostname()
	runtimeOs := strings.ReplaceAll(runtime.GOARCH, "amd", "x")
	directoryName := folderName + "-" + Os + "-" + runtimeOs
	zipFileName := directoryName + ".zip"

	executeCommand := exec.Command("./node_modules/.bin/nativefier", Url, "--name", folderName, "-p", Os)
	// executeCommand := exec.Command("nativefier", Url, "--name", folderName, "-p", Os)
	stdout, err := executeCommand.Output()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf(string(stdout))
	fmt.Printf("Zipping: %v\n", zipFileName)
	zipDirectory(directoryName, zipFileName)
	fmt.Printf("Zip complete! %v\n", zipFileName)
	return zipFileName

}

func zipDirectory(source, target string) error {
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

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
		defer f.Close()

		_, err = io.Copy(headerWriter, f)
		return err
	})
}
