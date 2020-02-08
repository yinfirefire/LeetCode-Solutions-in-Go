package main

type CombinationIterator struct {
	idx int
	ss  []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	ss := make([]string, 0)
	helper(0, "", characters, combinationLength, &ss)
	return CombinationIterator{0, ss}
}

func helper(idx int, temp, chars string, length int, ss *[]string) {
	if len(temp) == length {
		*ss = append(*ss, temp)
		return
	} else {
		for i := idx; i < len(chars); i++ {
			temp = temp + chars[i:i+1]
			helper(i+1, temp, chars, length, ss)
			temp = temp[0 : len(temp)-1]
		}
	}
}

func (this *CombinationIterator) Next() string {
	this.idx++
	return this.ss[this.idx]
}

func (this *CombinationIterator) HasNext() bool {
	return this.idx < len(this.ss)
}
