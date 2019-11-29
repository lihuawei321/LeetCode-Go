package main

//给定一个无向图graph，当这个图为二分图时返回true。
//如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。
//graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。每个节点都是一个在0到graph.length-1之间的整数。
//这图中没有自环和平行边： graph[i] 中不存在i，并且graph[i]中没有重复的值。

func isBipartite(graph [][]int) bool {
	v := len(graph)
	visited := make([]bool, v) //是否访问
	colors := make([]int, v)   //染色

	for i := 0; i < v; i++ {
		visited[i] = false
		colors[i] = 0
	}
	for u := 0; u < v; u++ {
		if !visited[u] {
			if !dfs(u, 0, visited, colors, graph) {
				return false
			}
		}
	}
	return true
}

func dfs(v, color int, visited []bool, colors []int, graph [][]int) bool {
	visited[v] = true
	colors[v] = color
	for _, w := range graph[v] {
		if !visited[w] {
			if !dfs(w, 1-color, visited, colors, graph) {
				return false
			}
		} else if colors[v] == colors[w] {
			return false
		}

	}
	return true
}

func main() {

}
