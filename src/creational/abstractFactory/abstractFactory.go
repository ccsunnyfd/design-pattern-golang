package main

import (
	"fmt"
	"runtime"
)

// button interface
type iButton interface {
	setSize(size int)
	getSize() int
	render()
}

// button struct
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
	fmt.Printf("html render. size: %d\n", hBtn.getSize())
}

// exact button
type windowsButton struct {
	button
}

func (wBtn *windowsButton) render() {
	fmt.Printf("windows render. size: %d\n", wBtn.getSize())
}

// checkbox interface
type iCheckbox interface {
	setRound(ifRound bool)
	getRound() bool
	render()
}

// checkbox struct
type checkbox struct {
	ifRound bool
}

func (checBox *checkbox) setRound(ifRound bool) {
	checBox.ifRound = ifRound
}

func (checBox *checkbox) getRound() bool {
	return checBox.ifRound
}

// exact checkbox
type htmlCheckbox struct {
	checkbox
}

func (hChecBox *htmlCheckbox) render() {
	fmt.Printf("html render. round: %t\n", hChecBox.getRound())
}

// exact checkbox
type windowsCheckbox struct {
	checkbox
}

func (wChecBox *windowsCheckbox) render() {
	fmt.Printf("windows render. round: %t\n", wChecBox.getRound())
}

// factory interface
type iGuiFactory interface {
	renderWindow()
	createButton() iButton
	createCheckbox() iCheckbox
}

// exact factory
type htmlFactory struct {
}

func (hFact *htmlFactory) createButton() iButton {
	return &htmlButton{
		button: button{
			size: 14,
		},
	}
}

func (hFact *htmlFactory) createCheckbox() iCheckbox {
	return &htmlCheckbox{
		checkbox: checkbox{
			ifRound: false,
		},
	}
}

func (hFact *htmlFactory) renderWindow() {
	hFact.createButton().render()
	hFact.createCheckbox().render()
}

// exact factory
type windowsFactory struct {
}

func (wFact *windowsFactory) createButton() iButton {
	return &windowsButton{
		button: button{
			size: 17,
		},
	}
}

func (wFact *windowsFactory) createCheckbox() iCheckbox {
	return &windowsCheckbox{
		checkbox: checkbox{
			ifRound: true,
		},
	}
}

func (wFact *windowsFactory) renderWindow() {
	wFact.createButton().render()
	wFact.createCheckbox().render()
}

// get factory factory
func getGuiFactory(osName string) (iGuiFactory, error) {
	if osName == "windows" {
		return &windowsFactory{}, nil
	}
	return &htmlFactory{}, nil
}

func main() {
	osName := runtime.GOOS
	factory, _ := getGuiFactory(osName)

	factory.renderWindow()
}
