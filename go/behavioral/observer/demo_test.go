package observer

import "testing"

func TestSubjectImpl_Notify(t *testing.T) {
	observer1 := Observer{1}
	observer2 := Observer{2}
	observer3 := Observer{3}

	event := &Event{Data: "new message coming!"}
	subjectImpl := SubjectImpl{Observers: make(map[Observer]interface{})}
	subjectImpl.Register(observer1)
	subjectImpl.Register(observer2)
	subjectImpl.Register(observer3)

	subjectImpl.Notify(event)

	//update:new message coming!,id:1
	//update:new message coming!,id:2
	//update:new message coming!,id:3
}
