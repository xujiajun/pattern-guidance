package builder

import (
	"testing"
	"fmt"
)

func TestGetCar(t *testing.T) {
	builder := BuilderCarImpl{}
	Director := NewDirector(&builder)

	Director.builder.BuildBody()
	Director.builder.BuildEngine()
	Director.builder.BuildWheel()

	newCar := Director.builder.BuildCar()

	fmt.Print(newCar)//{wheel1 engine1 body1}
}
