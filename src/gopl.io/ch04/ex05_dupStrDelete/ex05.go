package main

import "fmt"

func main() {
	data := []string{"one", "", "three","three","four","four","five"}
	fmt.Printf("%q\n", dupDelete(data))
	//fmt.Printf("%q\n", data)
}
func dupDelete(strings []string) []string {
	i := 0
	oldStr:="xxxxxxxxxxxxxxxxx"
	for _, s := range strings {
		if s != oldStr {
			strings[i] = s
			i++
		}
		oldStr=s
	}
	return strings[:i]
}
