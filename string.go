package main

import (
	"strconv"
	"strings"
)

// BM83-BM86 字符串

// BM83 字符串变形，小写大写呼唤，位置反转
func trans(s string, n int) string {
	// write code here
	mid := strings.Split(s, " ")
	for i := 0; i < len(mid)/2; i++ {
		mid[i], mid[len(mid)-1-i] = mid[len(mid)-1-i], mid[i]
	}
	res := []string{}
	for i := 0; i < len(mid); i++ {
		s := []byte(mid[i])
		for j := 0; j < len(s); j++ { //A的ASCII值为65，Z为90，a为97，相差32
			if s[j] >= 97 {
				s[j] -= 32
			} else {
				s[j] += 32
			}
		}
		res = append(res, string(s))
	}
	return strings.Join(res, " ")
}

// BM84 最长公共前缀
func longestCommonPrefix(strs []string) string {
	lstr := len(strs)
	if lstr == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < lstr; j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
func longestCommonPrefix1(strs []string) string {
	// write code here
	lstr := len(strs)
	if lstr == 1 {
		return strs[0]
	}
	if lstr == 0 {
		return ""
	}
	minlen := 5000
	for i := 0; i < lstr; i++ {
		if minlen > len(strs[i]) {
			minlen = len(strs[i])
		}
	}
	maxl := minlen
	for i := 1; i < lstr; i++ {
		num := 0
		for j := 0; j < minlen; j++ {
			if strs[0][j] == strs[i][j] {
				num += 1
			} else {
				break
			}
		}
		if num < maxl {
			maxl = num
		}
	}
	return strs[0][:maxl]
}

// BM85 验证IP地址
func solves(IP string) string {
	// write code here
	res := strings.Split(IP, ".")
	if len(res) != 1 {
		for i := 0; i < len(res); i++ {
			num, _ := strconv.Atoi(res[i])
			if num > 255 || num == 0 {
				return "Neither"
			}
			if res[i] == "" {
				return "Neither"
			}
			for _, s := range res[i] {
				if s >= '9' || s <= '0' {
					return "Neither"
				}
			}
		}
		return "IPv4"
	} else {
		res = strings.Split(IP, ":")
		for i := 0; i < len(res); i++ {
			if res[i] == "" {
				return "Neither"
			}
			for k, s := range res[i] {
				//每个“位”里第一个是数字0，直接下个循环
				if k == 0 && s == '0' {
					continue
				}
				//若第一个是0，第二个还是0，则不符合
				if k == 1 && s == '0' && res[i][0] == '0' {
					return "Neither"
				}
				//下列三种情况同时满足，表示字符不在合规的区间
				if (s < '0' || s > '9') && (s < 'a' || s > 'e') && (s < 'A' || s > 'E') {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
}
