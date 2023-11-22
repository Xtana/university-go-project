package Rooms

import (
	"fmt"
	"university-go-project/Home/Rooms/Device"
)

type Room struct {
	Name         string
	Width        float32
	Length       float32
	Height       float32
	WindowsCount int
	Device       Device.Devices
}

func (room Room) countSquare() float32 {
	return room.Width * room.Length
}

func (room Room) countPerimeter() float32 {
	return (room.Length + room.Width) * 2
}

func (room Room) countVolume() float32 {
	return room.Width * room.Length * room.Height
}

func (room Room) getRoomString() string {
	resString := ""
	resString += "\t\tКомната: " + room.Name + "\n" +
		"\t\tШирина: " + fmt.Sprint(room.Width) + "\n" +
		"\t\tДлинна: " + fmt.Sprint(room.Length) + "\n" +
		"\t\tВысота: " + fmt.Sprint(room.Height) + "\n" +
		"\t\tКоличество окон: " + fmt.Sprint(room.WindowsCount) + "\n"
	//"\t\t------------------------\n"
	return resString
}
