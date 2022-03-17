package main

import (
	"fmt"
	"gstream"
	"strconv"
)

// func Cast[T any](value T) T {
// 	return value
// }

func main() {
	st := gstream.NewStream([]int{1, 2, 3, 4}).MapToInt(gstream.Pass[int])
	s := gstream.Map(st.Stream, gstream.Pass[int]).Skip(1).ToSlice()
	s2 := gstream.NewStringStream([]string{"1", "2", "3"}).MapToInt(func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}).Sum()
	// s .Join(",")
	fmt.Println(s)
	fmt.Println(s2)
}
