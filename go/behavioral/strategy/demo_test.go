package strategy

import "testing"

func TestBook_GetMsg(t *testing.T) {
	study := &Study{name: "golang knowledge", strategy: &SE{}}
	study.strategy.GetMsg(study) //get msg golang knowledge from SE
}

func TestSE_GetMsg(t *testing.T) {
	study := &Study{name: "golang knowledge", strategy: &Book{}}
	study.strategy.GetMsg(study) //get msg golang knowledge from Book
}
