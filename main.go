package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"archive/zip"
	"io/ioutil"
	"bytes"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s stringtosearch", os.Args[0])
		os.Exit(1)
	}

	stringToSearch := bytes.NewBufferString(os.Args[1]).Bytes()

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if strings.Compare(".jar", ext) != 0 {
			return nil
		}

		jarFile, err := zip.OpenReader(path)
		if err != nil {
			return err
		}
		defer jarFile.Close()

		for _, f := range jarFile.File {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			content, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil
			}
			rc.Close()

			if bytes.Contains(content, stringToSearch) {
				fmt.Printf("found string %s in %s!%s", os.Args[1], path, f.Name)
			}
		}

		return nil
	})

}
