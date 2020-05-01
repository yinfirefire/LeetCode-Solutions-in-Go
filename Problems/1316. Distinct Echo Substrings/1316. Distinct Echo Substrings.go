package main

func distinctEchoSubstrings(text string) int {
	set := map[string]bool{}
	n := len(text)
	for w := 1; w <= n/2; w++ {
		cnt := 0
		for l, r := 0, w; r < n; l, r = l+1, r+1 {
			if text[l] == text[r] {
				cnt++
				if cnt == w {
					if !set[text[l-w+1:l+1]] {
						set[text[l-w+1:l+1]] = true
					}
				}
			} else {
				cnt = 0
			}
		}
	}
	return len(set)
}
