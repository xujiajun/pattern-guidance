package prototype

import "testing"

func TestResume_Clone(t *testing.T) {
	resume1 := &Resume{}
	resume1.SetPersonalInfo("johnny","male","28")
	resume1.SetWorkExperience("2012-2018","xxx company")

	resume2 := resume1.Clone()
	resume2.SetPersonalInfo("jack","male","29")
	resume2.SetWorkExperience("2013-2018","xxx2 company")

	resume1.Display()
	resume2.Display()

	//johnny male 28
	//Display： 2012-2018 xxx company
	//jack male 29
	//Display： 2013-2018 xxx2 company
}
