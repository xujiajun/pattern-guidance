package adapter

import (
	"log"
)

type Log interface {
	Info(string string)
	Warn(string string)
}

type Frame struct {
	log Log
}

type Alog struct {
}

func (alog *Alog) AInfo(s string) {
	log.Print(s)
}

func (alog *Alog) AWarn(s string) {
	log.Print(s)
}

type AdaptLog struct {
	alog *Alog
}

func (adaptLog *AdaptLog) Info(s string) {
	adaptLog.alog.AInfo(s)
}

func (adaptLog *AdaptLog) Warn(s string) {
	adaptLog.alog.AWarn(s)
}
