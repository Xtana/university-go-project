package home

import (
	"fmt"
	"university-go-project/home/family"
	"university-go-project/home/rooms"
	"university-go-project/home/rooms/device"
	"university-go-project/home/rooms/furniture"
	"university-go-project/home/rooms/furniture/things"
)

type Home struct {
	Family        family.Family
	Rooms         rooms.Rooms
	FamilySurname string
}

func (home Home) HomeInfo() {
	fmt.Println("\nДом семьи " + home.FamilySurname + ":")
	fmt.Println(home.Rooms.GeneralRoomParameters())
	fmt.Print(home.Family.FamilyInfo())
	fmt.Print(home.Rooms.RoomsInfo())
}

func Make() Home {
	i := family.FamilyMember{
		Name:               "Роман",
		Age:                20,
		Sex:                "М",
		RealAddress:        "Пр. Победы 00, 00",
		ResidentialAddress: "Крпемль 00, 28",
	}

	mam := family.FamilyMember{
		Name:               "Вероника",
		Age:                39,
		Sex:                "Ж",
		RealAddress:        "Пр. Победы 00, 00",
		ResidentialAddress: "Пр. Победы 00, 00",
	}

	dad := family.FamilyMember{
		Name:               "Андрей",
		Age:                43,
		Sex:                "М",
		RealAddress:        "Пр. Победы 00, 00",
		ResidentialAddress: "Революционная, 00",
	}

	livingRoomDevices := device.Devices{
		[]device.Device{
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

	bathroomDevices := device.Devices{
		[]device.Device{
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
	kitchenDevices := device.Devices{
		[]device.Device{
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

	hallwayDevices := device.Devices{
		[]device.Device{
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

	livingRoomFurnitureThingsTable := things.Things{
		[]things.Thing{
			{
				Type: "Фото",
			},
			{
				Type: "Писменные принадлежности",
			},
			{
				Type: "Тетрадь",
			},
		},
	}

	livingRoomFurnitureThingsSofa := things.Things{
		[]things.Thing{
			{
				Type: "Падушка",
			},
			{
				Type: "Покрывало",
			},
			{
				Type: "Плед",
			},
		},
	}

	livingRoomFurnitureThingsShelves := things.Things{
		[]things.Thing{
			{
				Type: "Иконы",
			},
			{
				Type: "Кникги",
			},
			{
				Type: "Утюг",
			},
			{
				Type: "Кубик Рубика",
			},
		},
	}

	livingRoomFurniture := furniture.Furnitures{
		[]furniture.Furniture{
			{
				Type:   "Письменный стол",
				Length: 250,
				Width:  130,
				Color:  "Brown",
				Things: livingRoomFurnitureThingsTable,
			},
			{
				Type:   "Кресло",
				Length: 70,
				Width:  100,
				Color:  "Black",
			},
			{
				Type:   "Стул",
				Length: 60,
				Width:  30,
				Color:  "Grey",
			},
			{
				Type:   "Диван",
				Length: 250,
				Width:  200,
				Color:  "White",
				Things: livingRoomFurnitureThingsSofa,
			},
			{
				Type:   "Полки",
				Length: 60,
				Width:  60,
				Color:  "Brown",
				Things: livingRoomFurnitureThingsShelves,
			},
		},
	}

	livingRoomFurnitureThingsKitchenArea := things.Things{
		[]things.Thing{
			{
				Type: "Доска",
			},
			{
				Type: "Полотенце",
			},
			{
				Type: "Солфетки",
			},
			{
				Type: "Тарелка с фруктами",
			},
		},
	}

	kitchenRoomFurniture := furniture.Furnitures{
		[]furniture.Furniture{
			{
				Type:   "Обеденный стол",
				Length: 300,
				Width:  80,
				Color:  "Black-white",
			},
			{
				Type:   "Диван",
				Length: 200,
				Width:  100,
				Color:  "White",
			},
			{
				Type:   "Стул",
				Length: 60,
				Width:  30,
				Color:  "Grey",
			},
			{
				Type:   "Кухонный уголок",
				Length: 250,
				Width:  200,
				Color:  "Beige",
				Things: livingRoomFurnitureThingsKitchenArea,
			},
		},
	}

	hallwayFurnitureBigCloset := things.Things{
		[]things.Thing{
			{
				Type: "Одежда",
			},
			{
				Type: "Горнолыжные принадлежности",
			},
		},
	}

	hallwayFurnitureSmallCloset := things.Things{
		[]things.Thing{
			{
				Type: "Всякое",
			},
		},
	}

	hallwayRoomFurniture := furniture.Furnitures{
		[]furniture.Furniture{
			{
				Type:   "Шкаф",
				Length: 200,
				Width:  70,
				Color:  "Black-brown",
				Things: hallwayFurnitureBigCloset,
			},
			{
				Type:   "Шкафчик",
				Length: 80,
				Width:  70,
				Color:  "Black-brown",
				Things: hallwayFurnitureSmallCloset,
			},
		},
	}

	livingRoom := rooms.Room{
		Name:         "Гостинная",
		Width:        3.5,
		Length:       5.5,
		Height:       2.5,
		WindowsCount: 1,
		Devices:      livingRoomDevices,
		Furnitures:   livingRoomFurniture,
	}

	kitchen := rooms.Room{
		Name:         "Кухня",
		Width:        2.3,
		Length:       2.5,
		Height:       2.5,
		WindowsCount: 1,
		Devices:      kitchenDevices,
		Furnitures:   kitchenRoomFurniture,
	}

	bathroom := rooms.Room{
		Name:         "Ванная комната",
		Width:        2.2,
		Length:       2,
		Height:       2.5,
		WindowsCount: 0,
		Devices:      bathroomDevices,
	}

	hallway := rooms.Room{
		Name:         "Коридор",
		Width:        3.5,
		Length:       1,
		Height:       2.5,
		WindowsCount: 0,
		Devices:      hallwayDevices,
		Furnitures:   hallwayRoomFurniture,
	}

	home := Home{
		FamilySurname: "Самохиных",
		Family:        family.Family{FamilyMemberArr: []family.FamilyMember{i, mam, dad}},
		Rooms:         rooms.Rooms{Rooms: []rooms.Room{livingRoom, kitchen, bathroom, hallway}},
	}

	return home
}
