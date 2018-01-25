package factory_method

import (
	"testing"
)

func createApi(factory ApiFactory, name string) Api {
	api := factory.create()
	api.SetName(name)
	return api
}

func TestSayApi_Say(t *testing.T) {

	var factory ApiFactory
	factory = SayApiFactory{}

	sayApi := createApi(factory, "SayApi")
	s := sayApi.Say("xujiajun")

	if (s != "Hi, xujiajun;name:SayApi") {
		t.Fatal("TestSayApi_Say test fail")
	}
}

func TestHelloApi_Say(t *testing.T) {

	var factory ApiFactory
	factory = HelloApiFactory{}

	helloApi := createApi(factory, "HelloApi")

	s := helloApi.Say("xujiajun")

	if (s != "Hi, xujiajun;name:HelloApi") {
		t.Fatal("TestHelloApi_Say test fail")
	}
}
