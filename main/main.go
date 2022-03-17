package main

import (
	"fmt"

	"github.com/face0u0/gstream"
)

func main() {

	gstream.NewStream([]*Bird{{"coco", 3}, {"pepper", 1}})
	gstream.NewStream([]int{1, -1, 5}).Filter(func(i int) bool {
		return i > 0
	}).ForEach(func(i int) { fmt.Println(i) })
	// st := gstream.NewStream([]int{3, 5, 1, 2, 4}).MapToInt(gstream.Pass[int]).Sort()
	// s, _ := gstream.Map(st.Stream, gstream.Pass[int]).Skip(1).ToSlice()
	// s2 := gstream.NewStringStream([]string{"1", "2", "3"}).MapToInt(func(s string) int {
	// 	i, _ := strconv.Atoi(s)
	// 	return i
	// }).Sum()
	// // s .Join(",")
	// fmt.Println(s)
	// fmt.Println(s2)
	// gstream.NewIntStream(Range(100)).ErrorFilter(func(i int) error {
	// 	return nil
	// })
	// gstream.NewIntStream(Range(100)).Parallel().ForEach(func(i int) {
	// 	fmt.Println(i)
	// })
}

type Bird struct {
	Name string
	Age  int
}

func Range(length int) []int {
	nl := make([]int, length)
	for i := 0; i < length; i++ {
		nl[i] = i
	}
	return nl
}
