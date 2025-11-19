package main

import "fmt"

type Device interface {
	Status() string
	TurnOn()
	TurnOff()
}

type Light struct {
	On bool
}

func (l *Light) Status() string {
	if l.On {
		return "Light is On"
	}
	return "Light is Off"
}

func (l *Light) TurnOn() {
	l.On = true
}

func (l *Light) TurnOff() {
	l.On = false
}

func main() {
	var smartLight Device
	smartLight = &Light{true}

	fmt.Println(smartLight.Status()) // Output: Light is On
	smartLight.TurnOff()
	fmt.Println(smartLight.Status()) // Output: Light is Off
}
