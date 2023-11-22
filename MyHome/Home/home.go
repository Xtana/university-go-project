package Home

import (
	"fmt"
	"university-go-project/Home/Family"
	"university-go-project/Home/Rooms"
	"university-go-project/Home/Rooms/Device"
)

type Home struct {
	Family        Family.Family
	Rooms         Rooms.Rooms
	FamilySurname string
}

func (home Home) HomeInfo() {
	fmt.Println("\nДом семьи " + home.FamilySurname + ":")
	fmt.Println(home.Rooms.GeneralRoomParameters())
	fmt.Println("\tЧлены семьи:")
	fmt.Print(home.Family.FamilyInfo())
	fmt.Println("\tКомнаты:")
	fmt.Print(home.Rooms.RoomsInfo())
}

func Make() Home {
	i := Family.FamilyMember{
		Name:               "Роман",
		Age:                20,
		Sex:                "М",
		RealAddress:        "Пр. Победы 18, 74",
		ResidentialAddress: "Залупкино 67, 87",
	}

	mam := Family.FamilyMember{
		Name:               "Вероника",
		Age:                39,
		Sex:                "Ж",
		RealAddress:        "Пр. Победы 18, 74",
		ResidentialAddress: "Пр. Победы 18, 74",
	}

	dad := Family.FamilyMember{
		Name:               "Андрей",
		Age:                43,
		Sex:                "М",
		RealAddress:        "Пр. Победы 18, 74",
		ResidentialAddress: "Революционная, 52",
	}

	livingRoomDevices := Device.Devices{
		[]Device.Device{
			{
				Type:    "Плазменный телевизор",
				Name:    "UP77 55'' 4K Smart UHD",
				Company: "LG",
				Length:  123.5,
				Width:   71.5,
				Color:   "Black",
				Count:   1,
			},

			{
				Type:    "Принтер",
				Name:    "M1120",
				Company: "Epson",
				Length:  41.5,
				Width:   22.15,
				Color:   "Black",
				Count:   1,
			},

			{
				Type:    "Акестическая система",
				Name:    "S-RS55TB",
				Company: "Pioneer",
				Length:  110.5,
				Width:   15.5,
				Color:   "Black",
				Count:   4,
			},

			{
				Type:    "Морозильная камера",
				Name:    "DFZ 5175",
				Company: "Indesit",
				Length:  190,
				Width:   75,
				Color:   "White",
				Count:   1,
			},
		},
	}

	bathroomDevices := Device.Devices{
		[]Device.Device{
			{
				Type:    "Стиральная машина",
				Name:    "F2J6NNFW",
				Company: "LG",
				Length:  60,
				Width:   50,
				Color:   "Black-White",
				Count:   1,
			},
		},
	}
	kitchenDevices := Device.Devices{
		[]Device.Device{
			{
				Type:    "Телевизор",
				Name:    "UE28H4000AKXRU",
				Company: "Samsung",
				Length:  65.5,
				Width:   37.15,
				Color:   "Black",
				Count:   1,
			},
			{
				Type:    "Электрическая варочная панель",
				Name:    "GIEV 613420 E",
				Company: "Grundig",
				Length:  58.5,
				Width:   58.5,
				Color:   "Black",
				Count:   1,
			},
			{
				Type:    "Микроволновка",
				Name:    "BF634RGS1",
				Company: "Siemens",
				Length:  60,
				Width:   45,
				Color:   "Grey",
				Count:   1,
			},
		},
	}

	hallwayDevices := Device.Devices{
		[]Device.Device{
			{
				Type:    "Робот-пылесос",
				Name:    "Legee-7",
				Company: "Hobot",
				Length:  33,
				Width:   34,
				Color:   "Blue-Black",
				Count:   1,
			},
		},
	}

	livingRoom := Rooms.Room{
		Name:         "Гостинная",
		Width:        3.5,
		Length:       5.5,
		Height:       2.5,
		WindowsCount: 1,
		Device:       livingRoomDevices,
	}

	kitchen := Rooms.Room{
		Name:         "Кухня",
		Width:        2.3,
		Length:       2.5,
		Height:       2.5,
		WindowsCount: 1,
		Device:       kitchenDevices,
	}

	bathroom := Rooms.Room{
		Name:         "Ванная комната",
		Width:        2.2,
		Length:       2,
		Height:       2.5,
		WindowsCount: 0,
		Device:       bathroomDevices,
	}

	hallway := Rooms.Room{
		Name:         "Коридор",
		Width:        3.5,
		Length:       1,
		Height:       2.5,
		WindowsCount: 0,
		Device:       hallwayDevices,
	}

	home := Home{
		FamilySurname: "Самохиных",
		Family:        Family.Family{FamilyMemberArr: []Family.FamilyMember{i, mam, dad}},
		Rooms: Rooms.Rooms{
			Rooms: []Rooms.Room{livingRoom, kitchen, bathroom, hallway},
		},
	}

	return home
}
