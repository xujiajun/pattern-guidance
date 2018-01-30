package observer

import (
	"fmt"
	"strconv"
)

type Observer struct {
	Id int
}

func (observer *Observer) update(event *Event) {
	fmt.Print("update:" + event.Data + ",id:" + strconv.Itoa(observer.Id) + "\n")
}

type Event struct {
	Data string
}

type Subject interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(event *Event)
}

type SubjectImpl struct {
	Observers map[Observer]interface{}
}

func (subjectImpl *SubjectImpl) Register(observer Observer) {
	subjectImpl.Observers[observer] = struct {
	}{}
}

func (subjectImpl *SubjectImpl) Remove(observer Observer) () {
	delete(subjectImpl.Observers, observer)
}

func (subjectImpl *SubjectImpl) Notify(event *Event) () {
	for ob, _ := range subjectImpl.Observers {
		ob.update(event)
	}
}
