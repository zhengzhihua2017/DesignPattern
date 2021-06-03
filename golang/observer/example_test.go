package observer

import "testing"

func TestObserver(t *testing.T) {
	s := NewSubject()
	pa := NewPlayer("a")
	pb := NewPlayer("b")
	s.Attach(pa)
	s.Attach(pb)
	s.SetStatus(1)
	s.SetStatus(2)
	s.SetStatus(3)
}
