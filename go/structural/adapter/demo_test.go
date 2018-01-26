package adapter

import (
	"testing"
)

func TestAdaptLog_Info(t *testing.T) {
	adaptLog := AdaptLog{&Alog{}}
	frame := Frame{log: &adaptLog}
	frame.log.Info("Info xxx")
}

func TestAdaptLog_Warn(t *testing.T) {
	adaptLog := AdaptLog{&Alog{}}
	frame := Frame{log: &adaptLog}
	frame.log.Warn("Warn xxx")
}
