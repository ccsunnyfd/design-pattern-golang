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

// exact button
type windowsButton struct {
	button
}

func (wBtn *windowsButton) render() {
	fmt.Printf("windows render. size: %d", wBtn.getSize())
}

type buttonFactory interface {
	renderWindow()
	createButton() iButton
}

// exact factory
type htmlButtonFactory struct {
}

func (hBtnFact *htmlButtonFactory) createButton() iButton {
	return &htmlButton{
		button: button{
			size: 14,
		},
	}
}

func (hBtnFact *htmlButtonFactory) renderWindow() {
	hBtnFact.createButton().render()
}

// exact factory
type windowsButtonFactory struct {
}

func (wBtnFact *windowsButtonFactory) createButton() iButton {
	return &windowsButton{
		button: button{
			size: 17,
		},
	}
}

func (wBtnFact *windowsButtonFactory) renderWindow() {
	wBtnFact.createButton().render()
}

func main() {
	if runtime.GOOS == "windows" {
		wBtnF := windowsButtonFactory{}
		wBtnF.renderWindow()
	} else {
		hBtnF := htmlButtonFactory{}
		hBtnF.renderWindow()
	}
}
