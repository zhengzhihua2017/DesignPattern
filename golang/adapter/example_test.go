package adapter

import "testing"

//TestAdapter
func TestAdapter(t *testing.T) {
	target := NewAdapter()
	target.Request()
}
