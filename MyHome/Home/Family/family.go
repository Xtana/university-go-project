package Family

type Family struct {
	FamilyMemberArr []FamilyMember
}

func (family Family) FamilyInfo() string {
	resString := ""
	for i, member := range family.FamilyMemberArr {
		resString += member.getFamilyMemberString()
		if (i + 1) != len(family.FamilyMemberArr) {
			resString += "\t\t------------------------\n"
		}
	}
	return resString
}
