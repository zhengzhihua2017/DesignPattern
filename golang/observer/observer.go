package observer

import "fmt"

type IObserver interface {
	Notify(int)
}

type player struct {
	name string
}

func NewPlayer(name string) *player {
	p := &player{
		name: name,
	}
	return p
}

func (p *player) Notify(status int) {
	fmt.Printf("%s %d\n", p.name, status)
}
