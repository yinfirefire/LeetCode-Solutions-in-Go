package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	set := make(map[string]bool)
	for _, s := range wordList {
		set[s] = true
	}
	visited := make(map[string]bool)
	queue := make([]string, 0, len(wordList))
	if beginWord == endWord {
		return 1
	}
	if !set[endWord] {
		return 0
	}
	queue = append(queue, beginWord)
	visited[beginWord] = true
	step := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur == endWord {
				return step
			}
			next := nextWords(cur, set)
			for _, n := range next {
				if !visited[n] {
					queue = append(queue, n)
					visited[n] = true
				}
			}
		}
		queue = queue[size:]
		step++
	}
	return 0
}

func nextWords(cur string, set map[string]bool) []string {
	var res []string
	for i := range cur {
		for c := 'a'; c <= 'z'; c++ {
			if cur[i] == byte(c) {
				continue
			}
			next := cur[:i] + string(c) + cur[i+1:]
			if set[next] {
				res = append(res, next)
			}
		}
	}
	return res
}

/*
The default type for character values is rune. That means, if you don’t declare a type explicitly when declaring a variable with a character value, then Go will infer the type as rune -

var firstLetter = 'A' // Type inferred as `rune` (Default type for character values)
You can create a byte variable by explicitly specifying the type -

var lastLetter byte = 'Z'
*/
