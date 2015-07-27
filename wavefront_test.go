package wavefront

import (
    "testing"
)

func TestRead(t *testing.T) {
	models, err := Read("gopher.obj")
	if err != nil {
		t.Fatalf("%v", err)
	}

	if len(models) != 12 {
		t.Fatalf("Expected number of objects to be %d", 12)
	}
}