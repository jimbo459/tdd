package walk

import "testing"

func TestWalk(t *testing.T) {
	expected := "James"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string){
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("Expected number of argument calls incorrect, expected %d, actual %d", 1, len(got))
	}

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}
