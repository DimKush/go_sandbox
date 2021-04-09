package datafile

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var deepLevel int = 0

func ReportPanic() {
	rec := recover()
	if rec == nil {
		return
	}

	err, ok := rec.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(rec)
	}

}

func TreeSubDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
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

			TreeSubDir(dirPath)

		} else {
			filesStr = append(filesStr, file.Name())
		}
	}

	for _, param := range filesStr {
		fmt.Printf("\n%s%s", levels.String(), param)
	}

	deepLevel = 0
}
