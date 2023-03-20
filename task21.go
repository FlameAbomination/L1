package main

import "fmt"

type Client struct {
}

type Device interface {
	ConnectMidi()
}

type MidiPort struct {
}

type USBPort struct {
}

type USBAdapter struct {
	USBPort *USBPort
}

func (w *USBAdapter) ConnectMidi() {
	fmt.Println("Adapter converts Midi signal to USB.")
	w.USBPort.ConnectUsb()
}
func (w *USBPort) ConnectUsb() {
	fmt.Println("USB connector is plugged into device.")
}

func (m *MidiPort) ConnectMidi() {
	fmt.Println("Midi connector is plugged device.")
}

func (c *Client) ConnectMidi(device Device) {
	fmt.Println("Client inserts Midi connector into device.")
	device.ConnectMidi()
}

func task21() {
	client := &Client{}
	midiDevice := &MidiPort{}

	client.ConnectMidi(midiDevice)

	USBDevice := &USBPort{}
	USBToMidiAdapter := &USBAdapter{
		USBPort: USBDevice,
	}

	client.ConnectMidi(USBToMidiAdapter)
}
