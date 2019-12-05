package main

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people, num_people)
	i := 1
	for candies > i {
		res[(i-1)%num_people] += i
		candies -= i
		i++
	}
	res[(i-1)%num_people] += candies
	return res
}
