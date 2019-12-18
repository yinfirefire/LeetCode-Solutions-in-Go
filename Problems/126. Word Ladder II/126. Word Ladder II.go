package main

var set map[string]bool
var graph map[string][]string
var res [][]string
var dist map[string]int

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	set = make(map[string]bool)
	graph = make(map[string][]string)
	dist = make(map[string]int)
	for _, s := range wordList {
		set[s] = true
	}
	set[beginWord] = true
	if !set[endWord] {

		return res
	}
	buildGraph()
	reverseBFS(beginWord, endWord)
	dfs(&[]string{}, beginWord, endWord)
	return res
}

func dfs(temp *[]string, cur string, target string) {
	*temp = append(*temp, cur)
	if cur == target {
		var copy []string
		copy = append(copy, *temp...)
		res = append(res, copy)
	} else {
		if next, ok := graph[cur]; ok {
			for _, neighbor := range next {
				if dist[neighbor] == dist[cur]-1 {
					dfs(temp, neighbor, target)
				}
			}
		}
	}
	*temp = (*temp)[:len(*temp)-1]
}

func reverseBFS(begin string, end string) {
	queue := make([]string, 0)
	queue = append(queue, end)
	visited := map[string]bool{}
	visited[end] = true
	step := 0
	dist[end] = 0
	for len(queue) > 0 {
		size := len(queue)
		step++
		for i := 0; i < size; i++ {
			cur := queue[i]
			for _, neighbor := range graph[cur] {
				if !visited[neighbor] {
					visited[neighbor] = true
					dist[neighbor] = step
					queue = append(queue, neighbor)
				}
			}
		}
		queue = queue[size:]
	}
}

func buildGraph() {
	for s := range set {
		next := make([]string, 0)
		for i := range s {
			for c := 'a'; c <= 'z'; c++ {
				if byte(c) != s[i] {
					neighbor := s[:i] + string(c) + s[i+1:]
					if set[neighbor] {
						next = append(next, neighbor)
					}
				}
			}
		}
		graph[s] = next
	}
}
