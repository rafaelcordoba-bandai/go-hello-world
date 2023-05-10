package go_hello_world

import "testing"

func TestGetHello(t *testing.T) {
	name := "John"
	got := GetHello(name)
	want := "Hello John"
	if got != want {
		t.Errorf("GetHello(%q) = %q, want %q", name, got, want)
	}
}
