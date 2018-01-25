package simple_factory

import "testing"

func TestHelloApi_Say(t *testing.T) {
	helloApi := NewApi(1)
	s := helloApi.Say("xujiajun")

	if s != "Hi, xujiajun" {
		t.Fatal("TestHelloApi_Say test fail")
	}

}
func TestHiApi_Say(t *testing.T) {
	helloApi := NewApi(2)
	s := helloApi.Say("xujiajun")

	if s != "Hello, xujiajun" {
		t.Fatal("TestHiApi_Say test fail")
	}
}
