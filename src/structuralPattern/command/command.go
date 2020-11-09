package main

import "fmt"

// iReceiver interface
type iReceiver interface {
	action(commandContent string)
}

// receiverA concrete receiver
type receiverA struct{}

func (r *receiverA) action(commandContent string) {
	fmt.Printf("ReceiverA执行了命令%s.......\n", commandContent)
}

// receiverB concrete receiver
type receiverB struct{}

func (r *receiverB) action(commandContent string) {
	fmt.Printf("ReceiverB执行了命令%s.......\n", commandContent)
}

// iCommand interface
type iCommand interface {
	execute()
}

// concreteCommandA
type concreteCommandA struct {
	r iReceiver
}

func (c *concreteCommandA) execute() {
	c.r.action("concreteCommandA")
}

// concreteCommandB
type concreteCommandB struct {
	r iReceiver
}

func (c *concreteCommandB) execute() {
	c.r.action("concreteCommandB")
}

// invoker
type invoker struct {
	c iCommand
}

func newInvoker(command iCommand) *invoker {
	return &invoker{
		c: command,
	}
}

func (i *invoker) call() {
	i.c.execute()
}

// main
func main() {
	rA := &receiverA{}
	cA := &concreteCommandA{
		r: rA,
	}
	invoker1 := &invoker{
		c: cA,
	}
	invoker1.call()

	rB := &receiverB{}
	cA = &concreteCommandA{
		r: rB,
	}
	invoker2 := &invoker{
		c: cA,
	}
	invoker2.call()

	cB := &concreteCommandB{
		r: rB,
	}
	invoker3 := &invoker{
		c: cB,
	}
	invoker3.call()
}
