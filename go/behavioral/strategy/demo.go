package strategy

import "fmt"

type Study struct {
	name     string
	strategy GetMsgStrategy
}

type GetMsgStrategy interface {
	GetMsg(studySubject *Study)
}

type SE struct {
}

func (se *SE) GetMsg(studySubject *Study) {
	fmt.Print("get msg " + studySubject.name + " from SE")
}

type Book struct {
}

func (se *Book) GetMsg(studySubject *Study) {
	fmt.Print("get msg " + studySubject.name + " from Book")
}
