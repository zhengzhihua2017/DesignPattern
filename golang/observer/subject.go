package observer

import "sync"

type Subject struct {
	arrayList []IObserver
	lock      sync.Mutex
}

func NewSubject() *Subject {
	s := &Subject{
		arrayList: make([]IObserver, 0),
	}
	return s
}

func (s *Subject) Attach(observer IObserver) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.arrayList = append(s.arrayList, observer)
}

func (s *Subject) SetStatus(status int) {
	s.notifyAllObservers(status)
}

func (s *Subject) notifyAllObservers(status int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, observer := range s.arrayList {
		observer.Notify(status)
	}
}
