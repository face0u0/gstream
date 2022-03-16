package gstream

type SampleInterfaceA interface {
	Hoge() string
}

type SampleStructA struct {
}

func (a *SampleStructA) Hoge() string {
	return "hoge"
}

type SampleStructB struct {
	SampleInterfaceA
}

func (b *SampleStructB) Hoge() string {
	return "hoge!!!"
}

// func main() {
// 	a := SampleStructA{}
// 	fmt.Println(a.Hoge()) // hoge
// 	b := SampleStructB{}
// 	fmt.Println(b.Hoge()) // hoge!!!
// }
