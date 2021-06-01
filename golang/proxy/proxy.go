package proxy

type proxy struct {
	sa *serverA
	sb *serverB
}

func NewProxy() *proxy {
	p := &proxy{
		sa: &serverA{},
		sb: &serverB{},
	}
	return p
}

func (p *proxy) Do(msg string) {
	p.sa.Do(msg)
}
