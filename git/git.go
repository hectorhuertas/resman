package git

type Git struct {
	Name string
}

type Local struct {
	Location string
}

func (l *Local) ID() string {
	return l.Location
}
