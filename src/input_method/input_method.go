package input_method

import (
	"fmt"
	"good_coder/trie"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type inputMethod struct {
	dictionary *trie.Trie
}

type singleCharacter struct {
	weight int
	character string
}

func New(dict []string) *inputMethod {
	im := new(inputMethod)
	im.dictionary = trie.NewTrie()

	for _, v := range dict {
		if strings.HasPrefix(v, "http") { // load from internet
			response, err := http.Get(v)
			if err != nil || response.StatusCode != http.StatusOK {
				fmt.Println("load dictionary failed:", err, ", path:", v)
				continue
			}

			body, _ := ioutil.ReadAll(response.Body)
			err = response.Body.Close()

			spell, ret := getSpell(v)
			if ret != 0 {
				fmt.Println("load dictionary failed:", err, ", path:", v)
				continue
			}

			im.parseDictionary(spell, string(body))
		} else { // load from file
			file, err := os.OpenFile(v, os.O_RDONLY, 0666)
			if err != nil {
				fmt.Println("load dictionary failed:", err, ", path:", v)
				continue
			}
			fileContent, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println("load dictionary failed:", err, ", path:", v)
				continue
			}

			spell, ret := getSpell(v)
			if ret != 0 {
				fmt.Println("load dictionary failed:", err, ", path:", v)
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