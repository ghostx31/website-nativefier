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

	folderName := name.Hostname()
	runtimeOs := strings.ReplaceAll(runtime.GOARCH, "amd", "x")
	directoryName := folderName + "-" + Os + "-" + runtimeOs
	zipFileName := directoryName + ".zip"

	executeCommand := exec.Command("nativefier", Url, "--name", folderName, "-p", Os)
	fmt.Println(zipFileName)

	stdout, err := executeCommand.Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(stdout))
	fmt.Printf("Zipping: %v", zipFileName)
	zipDirectory(directoryName, zipFileName)
	fmt.Println("Zip complete! ")
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
