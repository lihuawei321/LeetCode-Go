package main

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

func isAnagram(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}
	m := make(map[byte]int, ls)
	for i := 0; i < ls; i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}

	}
	return true

}
func main() {

}
