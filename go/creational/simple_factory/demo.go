package simple_factory

import "fmt"

type Api interface {
	Say(sth string) string
}

type HiApi struct {
}

type HelloApi struct {
}

func (*HiApi) Say(sth string) string {
	return fmt.Sprintf("Hi, %s", sth)
}

func (*HelloApi) Say(sth string) string {
	return fmt.Sprintf("Hello, %s", sth)
}

func NewApi(flag int) Api {
	if (flag == 1) {
		return &HiApi{}
	}

	if (flag == 2) {
		return &HelloApi{}
	}

	return nil
}
