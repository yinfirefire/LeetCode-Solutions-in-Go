package main

import "math"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}

//if there are two pigs, and we can test 3 times
/*
	1  2  3  4
	5  6  7  8
	9  10 11 12
	13 14 15 16
*/
//for the first time: pig one try [1,2,3,4], pig two try [1,5,9,13]
//for the second time: pig one try [5,6,7,8], pig two try [2,6,10,14]
//for the third time:...

//then based on the death, we can find the poisonous bucket (x,y) just like in a 2d cartesian coordinate
//If there are three pigs, we can have a 4*4*4 three-dimension space for the buckets
//the length of each axis is (test times+1)

//(no. of pigs)^(test times+1) >= no. of buckets
