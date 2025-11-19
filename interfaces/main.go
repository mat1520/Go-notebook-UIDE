package main

import "fmt"

type Device interface {
	Status() string
	TurnOn()
	TurnOff()
}

type Dimable interface {
	SetBrightness(level int)
}

type Speaker interface {
	Play(sound string)
}

type TemperatureControl struct {
	Temperature float64
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

func (l *Light) SetBrightness(level int) {
	fmt.Printf("Setting brightness to %d\n", level)
}

func main() {
	var smartLight Device
	smartLight = &Light{true}

	fmt.Println(smartLight.Status()) // Output: Light is On
	smartLight.TurnOn()
	fmt.Println(smartLight.Status()) // Output: Light is Off
	smartLight.TurnOff()

}
