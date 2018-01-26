package singleton

import (
	"testing"
	//"time"
	//"strconv"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetInstance("aa")
	instance2 := GetInstance("aa")

	if instance1 != instance2 {
		t.Fatal("instance is not equal")
	}

	//for i := 0; i < 10; i++ {
	//	go func() {
	//		s := GetInstance("get:" + strconv.Itoa(i))
	//		s.TestFunc()
	//	}()
	//}
	//time.Sleep(1e5)

}
