package datafile

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var deepLevel int = 0

func TreeSubDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Error during reading filesystem:", err)
		return err
	}

	var levels strings.Builder
	for i := 0; i < deepLevel*2; i++ {
		if i == 0 {
			levels.WriteString("|")
		}
		levels.WriteString("__")
	}

	var filesStr []string = []string{}

	for _, file := range files {
		if file.IsDir() {
			deepLevel++
			dirPath := filepath.Join(path, file.Name())
			fmt.Printf("\n%s%s/", levels.String(), file.Name())

			err := TreeSubDir(dirPath)
			if err != nil {
				log.Fatal("Error during reading filesystem:", err)
				return err
			}
		} else {
			filesStr = append(filesStr, file.Name())
		}
	}

	for _, param := range filesStr {
		fmt.Printf("\n%s%s", levels.String(), param)
	}

	deepLevel = 0
	return nil
}
