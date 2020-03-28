package main

import (
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	res := make([]string, 0)
	tb := make(map[string]map[int]string)
	for i := range transactions {
		ss := strings.Split(transactions[i], ",")
		name := ss[0]
		time, _ := strconv.Atoi(ss[1])
		location := ss[3]
		if _, ok := tb[name]; !ok {
			tb[name] = make(map[int]string)
		}
		if _, ok := tb[name][time]; ok {
			tb[name][time] = "xx"
		} else {
			tb[name][time] = location
		}
	}
	for i := range transactions {
		ss := strings.Split(transactions[i], ",")
		name := ss[0]
		time, _ := strconv.Atoi(ss[1])
		amount, _ := strconv.Atoi(ss[2])
		location := ss[3]
		if amount > 1000 {
			res = append(res, transactions[i])
			continue
		}
		v := tb[name]
		added := false
		for j := time - 60; j < time+61; j++ {
			if added {
				break
			}
			if _, ok := v[j]; ok && j != time && v[j] != location {
				added = true
				res = append(res, transactions[i])
			}
			if v[j] == "xx" && j == time {
				added = true
				res = append(res, transactions[i])
			}
		}
	}
	return res
}
