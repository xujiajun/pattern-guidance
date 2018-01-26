package singleton

import (
	"sync"
	"fmt"
)

type Instance interface {
	TestFunc()
}

type SingleTon struct {
	Attr string
}

func (this *SingleTon) TestFunc() {
	fmt.Println(this.Attr)
}

var instance *SingleTon
var once sync.Once

func GetInstance(str string) *SingleTon {
	once.Do(func() {
		instance = &SingleTon{Attr: str}
	})
	return instance
}
