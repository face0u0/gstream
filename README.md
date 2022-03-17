# gstream

gstream provides stream like api for slice.

## Requirements
`go +1.18`

## Getting Started

### Installing

To start using gstream, install Go and run `go get`
```
go get github.com/face0u0/gstream
```

### Try

gstream can handle each element one by one, using mehtod chain. In this case, `NewStream()` create stream object and each element is handled by `ForEach()`

```go
package main

import (
	"fmt"

	"github.com/face0u0/gstream"
)

func main() {
	gstream.NewStream([]string{"a", "b", "c"}).ForEach(func(s string) {
		fmt.Println(s)
	})
}
```
```
a
b
c
```

you can exec other operation before applying `ForEach()`

```go
func main() {
	gstream.NewStream([]int{1, -1, 5}).Filter(func(i int) bool {
		return i > 0
	}).ForEach(func(i int) { fmt.Println(i) })
}
```
```
1
5
```

## Stream 

1. create stream object from slice
2. apply intermediate operations
3. apply terminal operarion


### create stream object

```go
gstream.NewStream([]int{1, -1, 5})
```

```go
gstream.NewStream([]string{"a", "b", "c"})
```

stream can be also used for user defined type.
```go
type Bird struct {
	Name string
	Age  int
}
```
```go
gstream.NewStream([]*Bird{{"coco", 3}, {"pepper", 1}})
```

For some primitive type such as `int`, `float64` and `string` can be also created as below.

```go
gstream.NewIntStream([]int{1, -1, 5})
```
```go
gstream.NewStringStream([]string{"a", "b", "c"})
```

This enables a type peculiar operation such as `Sum()` for `int`.

## intermediate operations

### `Filter`
`Filter()` is  

## copyright
This software is released under the MIT License, see LICENSE.txt.