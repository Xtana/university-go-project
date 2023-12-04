package things

type Thing struct {
	Type string
}

func (thing Thing) getThingString() string {
	return "\t\t\t\t" + thing.Type + "\n"
}
