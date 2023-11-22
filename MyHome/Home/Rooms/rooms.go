package Rooms

import "fmt"

type Rooms struct {
	Rooms []Room
}

func (rooms Rooms) CountWindows() int {
	windowsSum := 0
	for _, room := range rooms.Rooms {
		windowsSum += room.WindowsCount
	}
	return windowsSum
}

func (rooms Rooms) CountSquare() float32 {
	var resSquare float32 = 0
	for _, room := range rooms.Rooms {
		resSquare += room.countSquare()
	}
	return resSquare
}

func (rooms Rooms) CountVolume() float32 {
	var resVolume float32 = 0
	for _, room := range rooms.Rooms {
		resVolume += room.countVolume()
	}
	return resVolume
}

func (rooms Rooms) RoomsInfo() string {
	resString := ""
	for _, room := range rooms.Rooms {
		resString += room.RoomInfo()
	}
	return resString
}

func (home Rooms) HomeInfo() string {
	resString := ""
	resString += "Площадь: " + fmt.Sprint(home.CountSquare()) + " м^2\n" +
		"Объем: " + fmt.Sprint(home.CountVolume()) + " м^3\n" +
		"Количество окон: " + fmt.Sprint(home.CountWindows())

	return resString
}
