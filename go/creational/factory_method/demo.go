package factory_method

import "fmt"

type Api interface {
	Say(sth string) string
	SetName(name string)
}

type ApiBase struct {
	name string
}

func (api *ApiBase) SetName(name string) {
	api.name = name
}

type SayApi struct {
	*ApiBase
}

type HelloApi struct {
	*ApiBase
}

type ApiFactory interface {
	create() Api
}

type SayApiFactory struct {
}

type HelloApiFactory struct {
}

//SayApiFactory是SayApi工厂类
func (SayApiFactory) create() Api {
	return &SayApi{&ApiBase{}}
}

//HelloApiFactory是HelloApi工厂类
func (HelloApiFactory) create() Api {
	return &HelloApi{&ApiBase{}}
}

func (api *SayApi) Say(sth string) string {
	return fmt.Sprintf("Hi, %s;name:%s", sth, api.ApiBase.name)
}

func (api *HelloApi) Say(sth string) string {
	return fmt.Sprintf("Hi, %s;name:%s", sth, api.ApiBase.name)
}
