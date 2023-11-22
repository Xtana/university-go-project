package Rooms

import "fmt"

type Room struct {
	Name         string
	Width        float32
	Length       float32
	Height       float32
	WindowsCount int
}

func (room Room) init() {
	room.countSquare()
	room.countVolume()
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

func (room Room) RoomInfo() string {
	resString := ""
	resString += "\tКомната: " + room.Name +
		"; Ширина: " + fmt.Sprint(room.Width) +
		"; Длинна: " + fmt.Sprint(room.Length) +
		"; Высота: " + fmt.Sprint(room.Height) +
		"; Количество окон: " + fmt.Sprint(room.WindowsCount) +
		"\n"
	return resString
}
