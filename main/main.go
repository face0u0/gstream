package main

import (
	"fmt"
	"gstream"
	"strconv"
)

// func Cast[T any](value T) T {
// 	return value
// }

func Range(length int) []int {
	nl := make([]int, length)
	for i := 0; i < length; i++ {
		nl[i] = i
	}
	return nl
}

func main() {
	st := gstream.NewStream([]int{3, 5, 1, 2, 4}).MapToInt(gstream.Pass[int]).Sort()
	s := gstream.Map(st.Stream, gstream.Pass[int]).Skip(1).ToSlice()
	s2 := gstream.NewStringStream([]string{"1", "2", "3"}).MapToInt(func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).Sum()
	// s .Join(",")
	fmt.Println(s)
	fmt.Println(s2)
	gstream.NewIntStream(Range(100)).Parallel().ForEach(func(i int) {
		fmt.Println(i)
	})
}
