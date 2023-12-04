package family

type Family struct {
	FamilyMemberArr []FamilyMember
}

func (family Family) FamilyInfo() string {
	resString := ""
	if len(family.FamilyMemberArr) > 0 {
		resString += "\tЧлены семьи:\n"
	}
	for i, member := range family.FamilyMemberArr {
		resString += member.getFamilyMemberString()
		if (i + 1) != len(family.FamilyMemberArr) {
			resString += "\t\t------------------------\n"
		}
	}
	return resString
}
