package facade

import "testing"

//TestRun
func TestRun(t *testing.T) {
	api := NewAPI()
	api.Run()
}
