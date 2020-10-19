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

type dialog interface {
	renderWindow()
	createButton() button
}

type htmlDialog struct {
	dialog
}

func (hDial *htmlDialog) createButton() button {
	return new(htmlButton)
}

func (hDial *htmlDialog) renderWindow() {
	hDial.createButton().render()
}

type windowsDialog struct {
	dialog
}

func (wDial *windowsDialog) createButton() button {
	return new(windowsButton)
}

func (wDial *windowsDialog) renderWindow() {
	wDial.createButton().render()
}

func main() {
	dialog := new(htmlDialog)
	dialog.renderWindow()
}
