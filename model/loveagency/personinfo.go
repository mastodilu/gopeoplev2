package loveagency

import "github.com/mastodilu/gopeoplev2/model/tools/smartphone"

// PersonInfo holds the information of a person
type PersonInfo struct {
	id   int
	sex  byte
	age  int
	chat chan<- *smartphone.Message
}

// ID getter
func (pi *PersonInfo) ID() int {
	return pi.id
}

// Sex getter
func (pi *PersonInfo) Sex() byte {
	return pi.sex
}

// Chat getter
func (pi *PersonInfo) Chat() chan<- *smartphone.Message {
	return pi.chat
}

// Age getter
func (pi *PersonInfo) Age() int {
	return pi.age
}

// NewPersonInfo creates a new PersonInfo type and returns its address
func NewPersonInfo(id int, sex byte, age int, chat chan<- *smartphone.Message) *PersonInfo {
	return &PersonInfo{
		id:   id,
		sex:  sex,
		age:  age,
		chat: chat,
	}
}
