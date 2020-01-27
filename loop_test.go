package tcploop

import "testing"

func TestLoop(t *testing.T) {
	want := "hello world!"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, wanted %q", got, want)
	}
}
