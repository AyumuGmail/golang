package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func main() {
	testValues := []int{1, 3, 19, 5, 2, 6, 21}
	Sort(testValues)
	fmt.Printf("%v\n", testValues)
}

//Sortはvaluesないの値をその中でソート
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

//appendValuesはtの要素をvaluesの正しい順序に追加して
//結果のスライスを返します。
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		//return &tree{value:value}と同じ
		fmt.Printf("DEBUG=%d\n", value)
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
