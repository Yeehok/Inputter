// Package input_method
/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/*
@file input_method.go
@author shenyihao(com@baidu.com)
@date 2021/11/25
@brief input engine
*/

package input_method

import (
	"inputter/trie"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type inputMethod struct {
	dictionary *trie.Trie
}

func New(dict []string) *inputMethod {
	im := new(inputMethod)
	im.dictionary = trie.NewTrie()

	for _, v := range dict {
		if strings.HasPrefix(v, "http") { // load from internet
			response, err := http.Get(v)
			if err != nil || response.StatusCode != http.StatusOK {
				continue
			}

			body, _ := ioutil.ReadAll(response.Body)
			err = response.Body.Close()

			spell, ret := getSpell(v)
			if ret != 0 {
				continue
			}

			im.parseDictionary(spell, string(body))
		} else { // load from file
			file, err := os.OpenFile(v, os.O_RDONLY, 0666)
			if err != nil {
				continue
			}
			fileContent, err := ioutil.ReadAll(file)
			if err != nil {
				continue
			}

			spell, ret := getSpell(v)
			if ret != 0 {
				continue
			}

			im.parseDictionary(spell, string(fileContent))
		}
	}

	return im
}

func getSpell(source string) (spell string, ret int) {
	endIndex := strings.LastIndex(source, ".dat")
	startIndex := strings.LastIndex(source, "/")

	if endIndex == -1 {
		return "", -1
	}

	if startIndex != -1 {
		return source[startIndex:endIndex], 0
	} else {
		return source[0: endIndex], 0
	}
}

func (im *inputMethod) parseDictionary(spell string, dict string) {
	im.dictionary.NewWord(spell, dict)
}

func (im *inputMethod) FindWords(spell string) (words []string) {
	return im.dictionary.FindWords(spell)
}