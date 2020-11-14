package main

import (
	"fmt"
	"sync"
)

// State
type State uint8

const (
	SMALL State = iota
	SUPER
	FIRE
	CAPE
)

func (s State) String() string {
	switch s {
	case SMALL:
		return "SMALL"
	case SUPER:
		return "SUPER"
	case FIRE:
		return "FIRE"
	case CAPE:
		return "CAPE"
	default:
		return "UNKNOWN"
	}
}

// iMario interface
type iMario interface {
	getName() State
	obtainMushRoom(*marioStateMachine)
	obtainCape(*marioStateMachine)
	obtainFireFlower(*marioStateMachine)
	meetMonster(*marioStateMachine)
}

// abstractMario
type abstractMario struct{}

func (a *abstractMario) getName() State {
	return SMALL
}

func (a *abstractMario) obtainMushRoom(m *marioStateMachine)   {}
func (a *abstractMario) obtainCape(m *marioStateMachine)       {}
func (a *abstractMario) obtainFireFlower(m *marioStateMachine) {}
func (a *abstractMario) meetMonster(m *marioStateMachine)      {}

// smallMarioOnce
var smallMarioOnce struct {
	sync.Once
	v *smallMario
}

// smallMario concrete mario
type smallMario struct {
	*abstractMario
}

func getSmallMarioInstance() *smallMario {
	smallMarioOnce.Do(
		func() {
			smallMarioOnce.v = &smallMario{}
		})
	return smallMarioOnce.v
}
func (s *smallMario) getName() State {
	return SMALL
}
func (s *smallMario) obtainMushRoom(m *marioStateMachine) {
	m.setCurrentState(getSuperMarioInstance())
	m.setScore(m.getScore() + 100)
}
func (s *smallMario) obtainCape(m *marioStateMachine) {
	m.setCurrentState(getCapeMarioInstance())
	m.setScore(m.getScore() + 200)
}
func (s *smallMario) obtainFireFlower(m *marioStateMachine) {
	m.setCurrentState(getFireMarioInstance())
	m.setScore(m.getScore() + 300)
}
func (s *smallMario) meetMonster(m *marioStateMachine) {}

// superMarioOnce
var superMarioOnce struct {
	sync.Once
	v *superMario
}

// superMario concrete mario
type superMario struct {
	am abstractMario
}

func getSuperMarioInstance() *superMario {
	smallMarioOnce.Do(
		func() {
			superMarioOnce.v = &superMario{}
		})
	return superMarioOnce.v
}
func (s *superMario) getName() State {
	return SUPER
}
func (s *superMario) obtainMushRoom(m *marioStateMachine) {}
func (s *superMario) obtainCape(m *marioStateMachine) {
	m.setCurrentState(getCapeMarioInstance())
	m.setScore(m.getScore() + 200)
}
func (s *superMario) obtainFireFlower(m *marioStateMachine) {
	m.setCurrentState(getFireMarioInstance())
	m.setScore(m.getScore() + 300)
}
func (s *superMario) meetMonster(m *marioStateMachine) {
	m.setCurrentState(getSmallMarioInstance())
	m.setScore(m.getScore() - 100)
}

// capeMarioOnce
var capeMarioOnce struct {
	sync.Once
	v *capeMario
}

// capeMario concrete mario
type capeMario struct {
	am abstractMario
}

func getCapeMarioInstance() *capeMario {
	capeMarioOnce.Do(
		func() {
			capeMarioOnce.v = &capeMario{}
		})
	return capeMarioOnce.v
}
func (s *capeMario) getName() State {
	return CAPE
}
func (s *capeMario) obtainMushRoom(m *marioStateMachine)   {}
func (s *capeMario) obtainCape(m *marioStateMachine)       {}
func (s *capeMario) obtainFireFlower(m *marioStateMachine) {}
func (s *capeMario) meetMonster(m *marioStateMachine) {
	m.setCurrentState(getSmallMarioInstance())
	m.setScore(m.getScore() - 200)
}

// fireMarioOnce
var fireMarioOnce struct {
	sync.Once
	v *fireMario
}

// fireMario concrete mario
type fireMario struct {
	am abstractMario
}

func getFireMarioInstance() *fireMario {
	fireMarioOnce.Do(
		func() {
			fireMarioOnce.v = &fireMario{}
		})
	return fireMarioOnce.v
}
func (s *fireMario) getName() State {
	return FIRE
}
func (s *fireMario) obtainMushRoom(m *marioStateMachine)   {}
func (s *fireMario) obtainCape(m *marioStateMachine)       {}
func (s *fireMario) obtainFireFlower(m *marioStateMachine) {}
func (s *fireMario) meetMonster(m *marioStateMachine) {
	m.setCurrentState(getSmallMarioInstance())
	m.setScore(m.getScore() - 300)
}

// marioStateMachine
type marioStateMachine struct {
	score        int
	currentState iMario
}

func newMarioStateMachine() *marioStateMachine {
	return &marioStateMachine{
		0,
		getSmallMarioInstance(),
	}
}

func (m *marioStateMachine) obtainMushRoom() {
	m.currentState.obtainMushRoom(m)
}

func (m *marioStateMachine) obtainCape() {
	m.currentState.obtainCape(m)
}

func (m *marioStateMachine) obtainFireFlower() {
	m.currentState.obtainFireFlower(m)
}

func (m *marioStateMachine) meetMonster() {
	m.currentState.meetMonster(m)
}

func (m *marioStateMachine) getScore() int {
	return m.score
}

func (m *marioStateMachine) getCurrentState() State {
	return m.currentState.getName()
}

func (m *marioStateMachine) setScore(score int) {
	m.score = score
}

func (m *marioStateMachine) setCurrentState(currentState iMario) {
	m.currentState = currentState
}

// main
func main() {
	mario := newMarioStateMachine()
	mario.obtainMushRoom()
	score := mario.getScore()
	state := mario.getCurrentState()
	fmt.Printf("mario score: %d; state: %s\n", score, state)

	mario.obtainFireFlower()
	score = mario.getScore()
	state = mario.getCurrentState()
	fmt.Printf("mario score: %d; state: %s\n", score, state)

	mario.meetMonster()
	score = mario.getScore()
	state = mario.getCurrentState()
	fmt.Printf("mario score: %d; state: %s\n", score, state)
}
