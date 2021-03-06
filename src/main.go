// Package main
/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/*
@file main.go
@author shenyihao(com@baidu.com)
@date 2021/11/25
@brief main
*/

package main

import (
	"bufio"
	"fmt"
	"inputter/input_method"
	"inputter/util"
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
			spell = strings.TrimRight(spell, util.WindowsEnter)
		} else {
			spell = strings.TrimRight(spell, util.UnixEnter)
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
