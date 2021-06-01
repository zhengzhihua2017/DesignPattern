package mediator

import (
	"testing"
)

func TestMediatort(t *testing.T) {
	m := NewMediator()
	m.Run(1)
	m.Run(2)
}
