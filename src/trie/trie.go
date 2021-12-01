// Package trie
/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/*
@file trie.go
@author shenyihao(com@baidu.com)
@date 2021/11/25
@brief trie implementation
*/

package trie

import (
	"container/list"
	"good_coder/util"
	"runtime"
	"strconv"
	"strings"
)

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return new(Trie).Init()
}

func (t *Trie) Init() *Trie {
	t.root = newNode()
	return t
}

// NewWord spell and dictionary
func (t *Trie) NewWord(spell string, dict string) {
	runeSpell := []byte(spell)

	parseDictionary := func(n *Node) {
		var line []string
		if runtime.GOOS == "windows" {
			line = strings.Split(dict, util.WindowsEnter)
		} else {
			line = strings.Split(dict, util.UnixEnter)
		}

		for _, v := range line {
			part := strings.Split(v, " ")
			if len(part) < 2 {
				continue
			}

			word := part[0]
			weight, err := strconv.Atoi(part[1])
			if err != nil {
				continue
			}

			if v, f := n.dictionary.Get(Int{weight}); f {
				arr := v.(*[]string)
				*arr = append(*arr, word)
			} else {
				arr := []string{word}
				n.dictionary.Insert(Int{weight}, &arr)
			}
		}
	}

	p := t.root
	l := len(spell)
	for i, v := range runeSpell {
		q := p.getChild(v)
		var vNode *Node = nil
		if q == nil {
			vNode = newNode()
			vNode.char = v
			p.child = append(p.child, vNode)
		} else {
			vNode = q
		}
		if i == l - 1 {
			parseDictionary(vNode)
		}
		p = vNode
	}
}

func (t *Trie) FindWords(spell string) (res []string) {
	byteSpell := []byte(spell)

	resMap := NewMap()

	// add words from n.dictionary
	insertWords := func(n *Node) {
		if n.dictionary.Len() <= 0 {
			return
		}
		for k, v := n.dictionary.Back(); k != nil && v != nil; k, v = n.dictionary.Previous(k) {
			s := *(v.(*[]string))
			if v2, f := resMap.Get(k); f {
				arr := v2.(*[]string)
				*arr = append(*arr, s...)
			} else {
				resMap.Insert(k, &s)
			}
		}
	}

	// Traversal using Breadth-First-Search
	var insertChild func(n *Node)
	insertChild = func(root *Node) {
		if root == nil {
			return
		}

		queue := list.New()
		queue.PushBack(root)

		for queue.Len() != 0 {
			n := queue.Front().Value.(*Node)

			insertWords(n)

			for _, v := range n.child {
				queue.PushBack(v)
			}
			queue.Remove(queue.Front())
		}
	}

	p := t.root
	l := len(byteSpell)
	for i, v := range byteSpell {
		q := p.getChild(v)
		if q == nil {
			return
		}

		p = q

		if i == l - 1 {
			insertChild(p)
		}
	}

	for k, v := resMap.Back(); k != nil && v != nil; k, v = resMap.Previous(k) {
		stringArr := *(v.(*[]string))
		res = append(res, stringArr...)
	}
	return
}

// Int key of trie
type Int struct {
	int
}

func (ik Int) Greater(v interface{}) bool {
	return ik.int > v.(Int).int
}

func (ik Int) Equal(v interface{}) bool {
	return ik.int == v.(Int).int
}
