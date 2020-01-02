package main

func canConvert(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	}
	mapping := make(map[byte]byte, 0)
	if len(str1) != len(str2) {
		return false
	}
	set := make(map[byte]bool, 0)
	for i := range str1 {
		if c, ok := mapping[str1[i]]; ok {
			if c != str2[i] {
				return false
			}
		} else {
			mapping[str1[i]] = str2[i]
			set[str2[i]] = true
		}
	}
	//note: we cannot just return true here
	//corner case:
	//"abcdefghijklmnopqrstuvwxyz"
	//"bcdefghijklmnopqrstuvwxyza"
	//we need to ensure that the total mapping characters appeared in str2 is at most 25
	//otherwise for the above example, we will run into a situation that
	//when we map the last character, there is another already transformed character in str1 to be changed
	//e.g. str1 = .....zz, str2 = .....za, the last step will transform two z at the end
	return len(set) <= 25
}

func main() {
	canConvert(
		"abcdefghijklmnopqrstuvwxyz",
		"bcdefghijklmnopqrstuvwxyzq")
}
