package Family

import (
	"fmt"
)

type Family struct {
	FamilyMembers []FamilyMember
}

func (family Family) FamilyInfo() string {
	resString := ""
	for _, member := range family.FamilyMembers {
		resString += "\tИмя: " + member.Name +
			"; Пол: " + member.Sex +
			"; Возраст: " + fmt.Sprint(member.Age) +
			"; Адрес проживания: " + member.RealAddress +
			"; Адрес прописки: " + member.ResidentialAddress +
			"\n"
	}
	return resString
}
