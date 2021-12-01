// Package trie
/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/*
@file node.go
@author shenyihao(com@baidu.com)
@date 2021/11/25
@brief trie node
*/

package trie

type Node struct {
	char byte
	dictionary *SortedMap
	child []*Node
}

func newNode() *Node {
	return new(Node).init()
}

func (n *Node) init() *Node {
	n.char = 0
	n.dictionary = NewMap()
	return n
}

// getChild Traversal children
func (n *Node) getChild(char byte) *Node {
	for _, v := range n.child {
		if v.char == char {
			return v
		}
	}
	return nil
}
