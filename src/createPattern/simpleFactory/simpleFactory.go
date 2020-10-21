package main

import (
	"fmt"
	"runtime"
)

type button interface {
	render()
}

type htmlButton struct {
	button
}

func (hbtn *htmlButton) render() {
	fmt.Println("htmlButton render...")
}

type windowsButton struct {
	button
}

func (wbtn *windowsButton) render() {
	fmt.Println("windowsButton render...")
}

type buttonFactory struct {
}

func (bf *buttonFactory) getButton(osName string) (button, error) {
	if "windows" == osName {
		return new(windowsButton), nil
	}
	return new(htmlButton), nil
}

func main() {
	newButton, _ := new(buttonFactory).getButton(runtime.GOOS)
	newButton.render()
}
