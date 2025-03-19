package greet

import "testing"

func TestGreet(t *testing.T) {
	substests := []struct {
		items  []string
		result string
	}{
		{
			items:  []string{"Tim"},
			result: "Hi, Tim!",
		},
		{
			items:  []string{},
			result: "Hi, world!",
		},
		{
			items:  []string{"Tim", "Jess"},
			result: "Hi, Tim, Jess!",
		},
	}

	for _, st := range substests {
		if s := Greet(st.items); s != st.result {
			t.Errorf("wanted %s [%v], got %s", st.items, st.result, s)
		}
	}
}
