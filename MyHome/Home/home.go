package Home

import (
	"fmt"
	"university-go-project/Home/Family"
)

type Home struct {
	Family        Family.Family
	FamilySurname string
}

func (home Home) HomeInfo() {
	fmt.Println("Дом семьи " + home.FamilySurname + "\n")
	fmt.Println("Члены семьи:")
	fmt.Print(home.Family.FamilyInfo())
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

	home := Home{Family: Family.Family{FamilyMembers: []Family.FamilyMember{i, mam, dad}},
		FamilySurname: "Самохиных",
	}

	return home
}
