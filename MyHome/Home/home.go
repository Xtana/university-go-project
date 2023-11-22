package Home

import (
	"fmt"
	"university-go-project/Home/Family"
	"university-go-project/Home/Rooms"
)

type Home struct {
	Family        Family.Family
	Rooms         Rooms.Rooms
	FamilySurname string
}

func (home Home) HomeInfo() {
	fmt.Println("Дом семьи " + home.FamilySurname)
	fmt.Println(home.Rooms.HomeInfo())
	fmt.Println("Члены семьи:")
	fmt.Print(home.Family.FamilyInfo())
	fmt.Println("Комнаты:")
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
		ResidentialAddress: "Пр. Победы 18, 74",
	}

	livingRoom := Rooms.Room{
		Name:         "Гостинная",
		Width:        3.5,
		Length:       5.5,
		Height:       2.5,
		WindowsCount: 1,
	}

	kitchen := Rooms.Room{
		Name:         "Кухня",
		Width:        2.3,
		Length:       2.5,
		Height:       2.5,
		WindowsCount: 1,
	}

	bathroom := Rooms.Room{
		Name:         "Ванная комната",
		Width:        2.2,
		Length:       2,
		Height:       2.5,
		WindowsCount: 0,
	}

	hallway := Rooms.Room{
		Name:         "Коридор",
		Width:        3.5,
		Length:       1,
		Height:       2.5,
		WindowsCount: 0,
	}

	home := Home{Family: Family.Family{FamilyMembers: []Family.FamilyMember{i, mam, dad}},
		FamilySurname: "Самохиных", Rooms: Rooms.Rooms{Rooms: []Rooms.Room{livingRoom, kitchen, bathroom, hallway}},
	}

	return home
}
