package Device

import "fmt"

type Device struct {
	Type    string
	Name    string
	Company string
	Length  float32
	Width   float32
	Color   string
	Count   int
}

func (devices Device) getDeviceString() string {
	resString := ""
	resString += "\t\t\tТип устройства: " + devices.Type + "\n" +
		"\t\t\tНазвание: " + fmt.Sprint(devices.Name) + "\n" +
		"\t\t\tФирма: " + fmt.Sprint(devices.Company) + "\n" +
		"\t\t\tДлинна: " + fmt.Sprint(devices.Length) + "\n" +
		"\t\t\tШирина: " + fmt.Sprint(devices.Width) + "\n" +
		"\t\t\tЦвет: " + fmt.Sprint(devices.Color) + "\n"
	//"\t\t\t------------------------\n"
	return resString
}
