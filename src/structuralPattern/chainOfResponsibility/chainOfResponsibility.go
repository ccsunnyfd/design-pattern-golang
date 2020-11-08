package main

import "fmt"

// handler interface
type handler interface {
	setSuccessor(handler)
	doHandle() bool
	handle(handler)
}

// baseHandler
type baseHandler struct {
	successor handler
}

func (b *baseHandler) setSuccessor(successor handler) {
	b.successor = successor
}

func (b *baseHandler) doHandle() bool {
	fmt.Println("抽象处理类。。。")
	return false
}

func (b *baseHandler) handle(self handler) {
	handled := self.doHandle()
	if b.successor != nil && !handled {
		b.successor.handle(b.successor)
	}
}

// handlerA concrete handler
type handlerA struct {
	baseHandler
}

func newHandlerA() *handlerA {
	hA := &handlerA{}
	return hA
}

func (hA *handlerA) doHandle() bool {
	handled := false
	fmt.Println("HandlerA is checking...")
	// handled = true
	return handled
}

// handlerB concrete handler
type handlerB struct {
	baseHandler
}

func newHandlerB() *handlerB {
	hB := &handlerB{}
	return hB
}

func (hB *handlerB) doHandle() bool {
	handled := false
	fmt.Println("HandlerB is checking...")
	handled = true
	return handled
}

// handlerChain
type handlerChain struct {
	head handler
	tail handler
}

func newHandlerChain() *handlerChain {
	return &handlerChain{
		head: nil,
		tail: nil,
	}
}

func (hC *handlerChain) addHandler(h handler) {
	h.setSuccessor(nil)

	if hC.head == nil {
		hC.head = h
		hC.tail = h
		return
	}

	hC.tail.setSuccessor(h)
	hC.tail = h
}

func (hC *handlerChain) handle() {
	if hC.head != nil {
		hC.head.handle(hC.head)
	}
}

// main
func main() {
	chain := newHandlerChain()
	chain.addHandler(newHandlerA())
	chain.addHandler(newHandlerB())
	chain.handle()
}
