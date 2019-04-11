package git

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
