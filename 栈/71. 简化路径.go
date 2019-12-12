package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	var res string
	var stack []string
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && len(cur) > 0 {
			stack = append(stack, cur)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return "/" + res
}
func main() {
	arr := []string{"1", "2", "3"}
	join := strings.Join(arr, "/")
	fmt.Println(join)
}
