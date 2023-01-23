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

func GetUrlFromUser(urlparams structs.Urlparams) string {
	fileName := BuildWebApp(urlparams)
	return fileName
}

func GetFilename(urlparams structs.Urlparams) (zipFileName string, folderName string, directoryName string) {
	name, err := url.Parse(urlparams.Url)
	if err != nil {
		panic(err)
	}
	fmt.Println(name.Hostname())

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
	fmt.Printf("\nZip file name is: %v", string(zipFileName))
	file, err := os.OpenFile("save.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write([]byte(zipFileName))
	if err != nil {
		panic(err)
	}
	return zipFileName, folderName, directoryName
}

func BuildWebApp(urlparams structs.Urlparams) string {
	zipFileName, folderName, directoryName := GetFilename(urlparams)
	// executeCommand := exec.Command("./node_modules/.bin/nativefier", urlparams.Url, "--name", folderName, "-p", urlparams.Os)
	executeCommand := exec.Command("nativefier", urlparams.Url, "--name", folderName, "-p", urlparams.Os)
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

// TODO: Will be used to read the save.txt file and provide the download to the user.
func userDownload() string {

	return ""
}
