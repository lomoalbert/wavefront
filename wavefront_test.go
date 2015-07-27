package wavefront

import (
    "testing"
    "fmt"
)

func TestRead(t *testing.T) {
	models, err := Read("gopher.obj")
	if err != nil {
		t.Fatalf("%v", err)
	}

	if len(models) != 12 {
		t.Fatalf("Expected number of objects to be %d", 12)
	}
    for key,obj := range models {
        fmt.Println(key,obj.Name)
        for _, group := range obj.Groups {
            fmt.Printf("\tmaterial name:%#v\n", group.Material.Name)
        }
    }
}