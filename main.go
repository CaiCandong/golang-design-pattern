package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(domain string) ([]string, int) {
	s := strings.Split(domain, " ")
	cnt, _ := strconv.Atoi(s[0])
	// 域名拆分
	s = strings.Split(s[1], ".")
	ret := []string{}
	for i := range s {
		ret = append(ret, strings.Join(s[i:], "."))
	}
	fmt.Println(ret, cnt)
	return ret, cnt
}
func main() {
	Parse("9001 discuss.leetcode.com")
}
