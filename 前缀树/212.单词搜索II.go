package main

import "reflect"

//给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
//同一个单元格内的字母在一个单词中不允许被重复使用。

func findWords2(board [][]byte, words []string) []string {
	res := []string{}
	for _, v := range words {
		if exist(board, v) {
			res = append(res, v)
		}
	}
	return res
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				if search(board, visited, word, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func search(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index >= len(word) {
		return true
	}
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return false
	}
	if visited[x][y] {
		return false
	}
	if board[x][y] != word[index] {
		return false
	}
	visited[x][y] = true
	if search(board, visited, word, index+1, x+1, y) || search(board, visited, word, index+1, x-1, y) || search(board, visited, word, index+1, x, y+1) || search(board, visited, word, index+1, x, y-1) {
		return true
	}
	visited[x][y] = false
	return false
}

//方法2

type Node struct {
	isWord   bool
	children map[rune]*Node
}

func (this *Node) insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.children[ch]; ok {
			parent = child
		} else {
			newChild := &Node{children: make(map[rune]*Node)}
			parent.children[ch] = newChild
			parent = newChild
		}
	}
	parent.isWord = true
}

func dfs(board [][]byte, path string, node Node, x, y int, res []string) bool {
	if node.isWord == true {
		res = append(res, path)
		node.isWord = false
	}
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return false
	}
	tmp := board[x][y]
	node = node.children
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	var trie Node
	for _, v := range words {
		trie.insert(v)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs()
		}
	}
	return res
}
func main() {

}
