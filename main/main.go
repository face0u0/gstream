package gstream

import (
	"fmt"
	"gstream"
)

func Cast[T any](value T) T {
	return value
}

func main() {
	s := gstream.NewStream([]int{1, 2, 3, 4}).MapToInt(Cast[int]).Sum()
	s2 := gstream.NewStream([]string{"a", "b", "c"}).MapToStr(Cast[string]).Join(",")
	// s .Join(",")
	fmt.Println(s)
	fmt.Println(s2)
}
