package proxy

import "fmt"

type serverA struct {
}

func (s *serverA) Do(msg string) {
	fmt.Printf("server a :%s\n", msg)
}

type serverB struct {
}

func (s *serverB) Do(msg string) {
	fmt.Printf("server b :%s\n", msg)
}
