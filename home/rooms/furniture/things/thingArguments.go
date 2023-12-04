package things

type Things struct {
	ThingArr []Thing
}

func (things Things) ThingsInfo() string {
	resString := ""
	if len(things.ThingArr) > 0 {
		resString += "\t\t\tВещи:\n"
	}
	for _, furn := range things.ThingArr {
		resString += furn.getThingString()
	}
	return resString
}
