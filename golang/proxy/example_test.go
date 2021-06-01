package proxy

import "testing"

func TestProxy(t *testing.T) {
	p := NewProxy()
	p.Do("hello")
}
