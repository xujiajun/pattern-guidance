package decorator

import (
	"testing"
	"fmt"
)

func TestDecorate(t *testing.T) {
	f := Decorate(Dressing)
	fmt.Println(f("coat"))//dressing coat and shoes
}
