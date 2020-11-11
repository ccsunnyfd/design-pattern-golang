package main

import "fmt"

// iTrain interface
type iTrain interface {
	arrive()
	depart()
	permitArrival()
}

// iMediator interface
type iMediator interface {
	canArrive(iTrain) bool
	notifyAboutDeparture()
}

// passengerTrain concrete train
type passengerTrain struct {
	m iMediator
}

func (p *passengerTrain) arrive() {
	if !p.m.canArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *passengerTrain) depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.m.notifyAboutDeparture()
}

func (p *passengerTrain) permitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.arrive()
}

// freightTrain concrete train
type freightTrain struct {
	m iMediator
}

func (f *freightTrain) arrive() {
	if !f.m.canArrive(f) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (f *freightTrain) depart() {
	fmt.Println("FreightTrain: Leaving")
	f.m.notifyAboutDeparture()
}

func (f *freightTrain) permitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	f.arrive()
}

// stationManager concrete mediator
type stationManager struct {
	isPlatformFree bool
	trainQueue     []iTrain
}

func newStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

func (s *stationManager) canArrive(t iTrain) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

// main
func main() {
	stationManager := newStationManager()

	passengerTrain := &passengerTrain{
		m: stationManager,
	}
	freightTrain := &freightTrain{
		m: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
