package main

import "fmt"

// message
type message struct {
	text string
}

func newMessage(m string) *message {
	return &message{
		m,
	}
}

func (m *message) getText() string {
	return m.text
}

// subject interface
type subject interface {
	registerObserver(o observer)
	removeObserver(o observer)
	notifyAll(m *message)
}

// observer interface
type observer interface {
	update(m *message)
}

// concreteSubject
type concreteSubject struct {
	observers []observer
}

func newConcreteSubject() *concreteSubject {
	return &concreteSubject{}
}

func (s *concreteSubject) registerObserver(o observer) {
	s.observers = append(s.observers, o)
}

func (s *concreteSubject) removeObserver(o observer) {
	s.observers = removeFromSlice(s.observers, o)
}

func removeFromSlice(observers []observer, observerToRemove observer) []observer {
	observerListLength := len(observers)
	for i, observer := range observers {
		if observerToRemove == observer {
			observers[observerListLength-1], observers[i] = observers[i], observers[observerListLength-1]
			return observers[:observerListLength-1]
		}
	}
	return observers
}

func (s *concreteSubject) notifyAll(m *message) {
	for _, o := range s.observers {
		o.update(m)
	}
}

// concreteObserverOne
type concreteObserverOne struct{}

func (o *concreteObserverOne) update(m *message) {
	fmt.Printf("ConcreteObserverOne is notified. %s\n", m.getText())
}

// concreteObserverTwo
type concreteObserverTwo struct{}

func (o *concreteObserverTwo) update(m *message) {
	fmt.Printf("ConcreteObserverTwo is notified. %s\n", m.getText())
}

// main
func main() {
	s := newConcreteSubject()
	o1 := &concreteObserverOne{}
	o2 := &concreteObserverTwo{}
	s.registerObserver(o1)
	s.registerObserver(o2)
	s.notifyAll(newMessage("A dog lingered!"))
	s.removeObserver(o2)
	s.notifyAll(newMessage("A cat lingered!"))
}
