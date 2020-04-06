package main

import (
	"fmt"
	"sort"
)

func sortDemo1() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

// 自定义类型
type byLength []string

// 实现sort.Interface 的三个方法
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 自定义规则排序
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func sortDemo2() {
	fruits := []string{"peach", "banana", "kiwi"}
	// 类型强制转换
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
