package main

//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/word-search

var dir = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if search(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func search(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if board[x][y] == word[index] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && search(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		visited[x][y] = false
	}
	return false

}

//方法2

func exist2(board [][]byte, word string) bool {
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

func search2(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
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
func main() {

}
