package prototype

import "fmt"

type PersonalInfo struct {
	name string
	sex  string
	age  string
}

type WorkExperience struct {
	timeArea string
	company  string
}

type Resume struct {
	PersonalInfo
	WorkExperience
}

func (this *Resume) SetPersonalInfo(name string, sex string, age string) {
	this.name = name
	this.sex = sex
	this.age = age
}

func (this *Resume) SetWorkExperience(timeArea string, company string) {
	this.timeArea = timeArea
	this.company = company
}

func (this *Resume) Display() {
	fmt.Println(this.name, this.sex, this.age)
	fmt.Println("Display：", this.timeArea, this.company)
}

func (this *Resume) Clone() *Resume {
	resume := *this
	return &resume
}
