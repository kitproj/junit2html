package example

import "testing"

func TestPassing(t *testing.T) {
	ok()
}

func TestFailing(t *testing.T) {
	t.Errorf("this test failed")
}

func TestSkipped(t *testing.T) {
	t.Skip("this test is skipped")
}

func TestPanic(t *testing.T) {
	kaboom()
}
