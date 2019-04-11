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

type LocalStore interface {
	Adder
	Getter
	Deleter
}

type Adder interface {
	Add(Local) error
}

type Getter interface {
	Get(string) (Local, error)
}
type Deleter interface {
	Delete(string)
}
