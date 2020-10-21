package main

import "fmt"

type button interface {
	render()
	onClick()
}

type htmlButton struct {
	button
}

func (hbtn *htmlButton) render() {
	fmt.Println("htmlButton render...")
	hbtn.onClick()
}

func (hbtn *htmlButton) onClick() {
	fmt.Println("htmlButton onClick...")
}

type windowsButton struct {
	button
}

func (wbtn *windowsButton) render() {
	fmt.Println("windowsButton render...")
	wbtn.onClick()
}

func (wbtn *windowsButton) onClick() {
	fmt.Println("windowsButton onClick...")
}

type buttonFactory interface {
	renderWindow()
	createButton() button
}

type htmlButtonFactory struct {
	buttonFactory
}

func (hDial *htmlButtonFactory) createButton() button {
	return new(htmlButton)
}

func (hDial *htmlButtonFactory) renderWindow() {
	hDial.createButton().render()
}

type windowsButtonFactory struct {
	buttonFactory
}

func (wDial *windowsButtonFactory) createButton() button {
	return new(windowsButton)
}

func (wDial *windowsButtonFactory) renderWindow() {
	wDial.createButton().render()
}

func main() {
	dialog := new(windowsButtonFactory)
	dialog.renderWindow()
}
