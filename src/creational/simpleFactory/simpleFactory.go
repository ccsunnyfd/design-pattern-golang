package main

import (
	"fmt"
	"runtime"
)

type iButton interface {
	setSize(size int)
	getSize() int
	render()
}

type button struct {
	size int
}

func (btn *button) setSize(size int) {
	btn.size = size
}

func (btn *button) getSize() int {
	return btn.size
}

// exact button
type htmlButton struct {
	button
}

func (hBtn *htmlButton) render() {
	fmt.Printf("html render. size: %d", hBtn.getSize())
}

func newHTMLButton() iButton {
	return &htmlButton{
		button: button{
			size: 14,
		},
	}
}

// exact button
type windowsButton struct {
	button
}

func (wBtn *windowsButton) render() {
	fmt.Printf("windows render. size: %d", wBtn.getSize())
}

func newWindowsButton() iButton {
	return &windowsButton{
		button: button{
			size: 17,
		},
	}
}

// getButtonFactory
func getButtonFactory(osName string) (iButton, error) {
	if osName == "windows" {
		return newWindowsButton(), nil
	}
	return newHTMLButton(), nil
}

func main() {
	newButton, _ := getButtonFactory(runtime.GOOS)
	newButton.render()
}
