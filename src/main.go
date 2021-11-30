package main

import (
	"bufio"
	"fmt"
	"good_coder/input_method"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func loop(files []string) {
	im := input_method.New(files)
	stdin := bufio.NewReader(os.Stdin)
	for {
		spell, err := stdin.ReadString('\n')
		if err != nil {
			break
		}
		if runtime.GOOS == "windows" {
			spell = strings.TrimRight(spell, "\n\r")
		} else {
			spell = strings.TrimRight(spell, "\n")
		}
		words := im.FindWords(spell)
		fmt.Println(strings.Join(words, ", "))
	}
}

func listDir(dirPth string) (file []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {

		if !fi.IsDir() {
			file = append(file, fi.Name())
		}
	}
	return file, nil
}

func main() {
	if len(os.Args) == 1 {
		files, _ := listDir(".")
		loop(files)
	} else {
		loop(os.Args[1:])
	}
}
