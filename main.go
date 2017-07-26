package main

import "fmt"
import "strconv"
import "sort"

func ReturnInt() int {
	i := 1
	return i
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(a []int) string {
	var s string
	for i := range a {
		s += strconv.Itoa(a[i])
	}
	return s
}

func MergeSlices(a []float32, b []int32) []int {
	c := make([]int, len(a))
	for i := range a {
		c[i] = int(a[i])
	}
	for i := range b {
		c = append(c, int(b[i]))
	}
	return c
}

func GetMapValuesSortedByKey(a map[int]string) []string {
	var keys []int
	var values []string
	for k := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		values = append(values, a[k])
	}
	return values
}

func main() {
	fmt.Println("все гуд")
	fmt.Println("все гуд еще раз")
}
