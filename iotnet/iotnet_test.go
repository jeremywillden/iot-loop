package iotnet

import "testing"

func TestLoop(t *testing.T) {
	want := "pi-cluster-00"
	if got := Myname(); got != want {
		t.Errorf("Myname() = %q, wanted %q", got, want)
	}
}
