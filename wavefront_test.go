package wavefront

import (
    "testing"
    "fmt"
)

func TestRead(t *testing.T) {
	model, err := Read("gopher.obj")
	if err != nil {
		t.Fatalf("%v", err)
	}

	if len(model) != 12 {
		t.Fatalf("Expected number of objects to be %d", 12)
	}
    fmt.Printf("%#v\n",model["Hnad_L_Sphere.012"].Groups[0].Material)
}
